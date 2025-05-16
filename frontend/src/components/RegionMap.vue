<template>
  <v-chart
    v-if="mapLoaded"
    class="h-full w-full"
    :option="chartOption"
    autoresize
    @click="handleMapClick"
  />
  <div v-else class="loading">Loading map data...</div>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted, watch } from 'vue'
import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import {
  GeoComponent,
  TooltipComponent,
  VisualMapComponent,
  LegendComponent
} from 'echarts/components'
import { MapChart } from 'echarts/charts'
import VChart from 'vue-echarts'
import * as echarts from 'echarts'
import { useSessionStore } from '@/stores/session'
import type { Region } from '@/tables'

// Register necessary ECharts components
use([
  CanvasRenderer,
  GeoComponent,
  MapChart,
  TooltipComponent,
  VisualMapComponent,
  LegendComponent
])

// Define the marker data structure
export interface MapMarker {
  name: string
  coordinate: [number, number] // [longitude, latitude]
  itemStyle?: {
    color?: string
  }
  symbol?: string
  symbolSize?: number
  tooltip?: {
    formatter?: string | Function
  }
}

const props = withDefaults(defineProps<{
  title?: string
  markers?: MapMarker[]
  showRegionColors?: boolean
  colorProperty?: string
  minValue?: number | null
  maxValue?: number | null
  colorRange?: string[]
  zoom?: number
  center?: [number, number] | null
  height?: string
  series?: {
    name: string,
    value: number,
  }[]
}>(), {
  title: '',
  markers: () => [],
  showRegionColors: false,
  colorProperty: 'forecast',
  minValue: null,
  maxValue: null,
  colorRange: () => ['#a5f3fd', '#58ccdc', '#119fc7', '#0b5b79', '#04202f'],
  zoom: 1,
  center: null,
  height: '600px'
})

// Define events
const emit = defineEmits<{
  (e: 'regionClick', region: any): void
  (e: 'markerClick', marker: any): void
  (e: 'mapLoaded', geoJson: any): void
}>()

// Component state
//const chartContainer = ref<HTMLElement | null>(null)
const sessionStore = useSessionStore()
const mapLoaded = ref(false)
const mapGeoJson = ref<any>(null)

// Convert our regions data to GeoJSON format
const convertRegionsToGeoJSON = (regions: Region[]) => {
  const features = regions.map(region => {
    return {
      type: 'Feature',
      properties: {
        id: region.id,
        name: region.name,
        description: region.description,
        altitude: region.altitude,
        drainage: region.drainage,
        forecast: region.forecast
      },
      geometry: {
        type: 'MultiPolygon',
        coordinates: region.coordinate
      }
    }
  })

  return {
    type: 'FeatureCollection',
    features
  }
}

// Compute min and max values for visualMap
const valueRange = computed(() => {
  if (props.minValue !== null && props.maxValue !== null) {
    return [props.minValue, props.maxValue]
  }

  if (!sessionStore.regions.length) return [0, 100]

  const values = sessionStore.regions
    .map(region => region[props.colorProperty as keyof Region])
    .filter(val => val !== null) as number[]

  if (values.length === 0) return [0, 100]

  return [
    Math.min(...values),
    Math.max(...values)
  ]
})

// Chart options computed property
const chartOption = computed(() => {
  if (!mapGeoJson.value) return {}

  const option: any = {
    title: props.title ? {
      text: props.title,
      left: 'center'
    } : undefined,
    tooltip: {
      trigger: 'item',
      formatter: (params: any) => {
        if (params.seriesType === 'map') {
          const properties = params.data?.properties || {}
          return `
            <div>
              <strong>${properties.name || 'Unknown'}</strong><br/>
              ${properties.description ? `<div>${properties.description}</div>` : ''}
              ${properties.altitude !== null ? `<div>区域海拔指数: ${properties.altitude}</div>` : ''}
              ${properties.drainage !== null ? `<div>排水完善指数: ${properties.drainage}</div>` : ''}
              ${properties.forecast !== null ? `<div>预报雨量指数: ${properties.forecast}</div>` : ''}
            </div>
          `
        } else if (params.seriesType === 'scatter') {
          return `<div><strong>${params.name}</strong></div>`
        }
        return ''
      }
    },
    geo: {
      map: 'floodMap',
      roam: true,
      zoom: props.zoom,
      center: props.center,
      itemStyle: {
        areaColor: '#f3f3f3',
        borderColor: '#666',
        borderWidth: 0.5
      },
      emphasis: {
        itemStyle: {
          areaColor: '#cce5ff'
        },
        label: {
          show: true
        }
      }
    },
    series: props.series ?? [],
  }

  // Add map series
  const mapSeries = {
    name: 'Regions',
    type: 'map',
    map: 'floodMap',
    geoIndex: 0,
    data: mapGeoJson.value.features.map((feature: any) => ({
      name: feature.properties.name,
      value: feature.properties[props.colorProperty] || 0,
      properties: feature.properties
    })),
    roam: true,
    zoom: props.zoom,
    center: props.center,
    emphasis: {
      label: {
        show: true
      }
    }
  }
  option.series.push(mapSeries)

  // Add visual map if region coloring is enabled
  if (props.showRegionColors) {
    option.visualMap = {
      type: 'continuous',
      min: valueRange.value[0],
      max: valueRange.value[1],
      text: [`Max: ${valueRange.value[1]}`, `Min: ${valueRange.value[0]}`],
      inRange: {
        color: props.colorRange
      },
      calculable: true,
      orient: 'horizontal',
      left: 'center',
      bottom: '5%'
    }
  }

  // Add markers if enabled
  if (props.markers.length > 0) {
    option.series.push({
      name: 'Markers',
      type: 'scatter',
      coordinateSystem: 'geo',
      data: props.markers.map(marker => ({
        name: marker.name,
        value: marker.coordinate,
        itemStyle: marker.itemStyle,
        symbol: marker.symbol || 'pin',
        symbolSize: marker.symbolSize || 20
      })),
      label: {
        formatter: '{b}',
        position: 'right',
        show: true
      }
    })
  }

  return option
})

// Handle map clicks
const handleMapClick = (params: any) => {
  if (params.seriesType === 'map') {
    emit('regionClick', params.data?.properties)
  } else if (params.seriesType === 'scatter') {
    emit('markerClick', params.data)
  }
}

// Load map data
const loadMapData = async () => {
  // Check if map data is already loaded
  if (sessionStore.regions.length === 0) {
    try {
      await sessionStore.loadMap()
    } catch (error) {
      console.error('Failed to load map data:', error)
      return
    }
  }

  // Convert regions data to GeoJSON
  mapGeoJson.value = convertRegionsToGeoJSON(sessionStore.regions)

  // Register map with ECharts
  echarts.registerMap('floodMap', mapGeoJson.value)

  mapLoaded.value = true
  emit('mapLoaded', mapGeoJson.value)
}

// Watch for changes in map data
watch(() => sessionStore.regions, (newMap) => {
  if (newMap.length > 0) {
    mapGeoJson.value = convertRegionsToGeoJSON(newMap)
    echarts.registerMap('floodMap', mapGeoJson.value)
    mapLoaded.value = true
  }
}, { deep: true })

// Initialize component
onMounted(async () => {
  await loadMapData()
})
</script>

<style scoped>

.loading {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  width: 100%;
  font-size: 16px;
  color: #666;
}
</style>
