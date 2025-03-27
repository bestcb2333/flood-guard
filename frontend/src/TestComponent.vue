<template>
  <div class="map-container">
    <v-chart class="chart" :option="chartOption" autoresize @click="handleMapClick" />
  </div>
</template>

<script setup lang="ts">
import { reactive, onMounted, computed } from 'vue';
import { use } from 'echarts/core';
import { CanvasRenderer } from 'echarts/renderers';
import { GeoComponent, TooltipComponent, VisualMapComponent, LegendComponent } from 'echarts/components';
import { MapChart, ScatterChart, EffectScatterChart } from 'echarts/charts';
import VChart from 'vue-echarts';
import * as echarts from 'echarts';

// 注册必要的组件
use([
  CanvasRenderer,
  GeoComponent,
  MapChart,
  ScatterChart,
  EffectScatterChart,
  TooltipComponent,
  VisualMapComponent,
  LegendComponent
]);

// 地图状态
const mapState = reactive({
  selectedRegion: '',
  hoveredRegion: '',
  points: [
    { name: '点位1', value: [120.38, 37.35, 10], selected: false },
    { name: '点位2', value: [122.207216, 37.433525, 15], selected: false },
    { name: '点位3', value: [121.15, 31.89, 8], selected: false },
    { name: '点位4', value: [120.33, 36.07, 20], selected: false },
  ]
});

// 注册地图
onMounted(() => {
  echarts.registerMap('customMap', {
  "type": "FeatureCollection",
  "features": [
    {
      "type": "Feature",
      "properties": {
        "name": "区域A",
        "id": "001"
      },
      "geometry": {
        "type": "Polygon",
        "coordinates": [[[120, 30], [125, 30], [125, 35], [120, 35], [120, 30]]]
      }
    },
    {
      "type": "Feature",
      "properties": {
        "name": "区域B",
        "id": "002"
      },
      "geometry": {
        "type": "Polygon",
        "coordinates": [[[120, 35], [125, 35], [125, 40], [120, 40], [120, 35]]]
      }
    },
    {
      "type": "Feature",
      "properties": {
        "name": "区域C",
        "id": "003"
      },
      "geometry": {
        "type": "Polygon",
        "coordinates": [[[115, 35], [120, 35], [120, 40], [115, 40], [115, 35]]]
      }
    }
  ]
});
});

// 处理地图点击事件
const handleMapClick = (params: any) => {
  if (params.componentType === 'series') {
    if (params.seriesType === 'map') {
      // 点击多边形区域
      mapState.selectedRegion = params.name;
      console.log('选中区域:', params.name);
    } else if (params.seriesType === 'effectScatter') {
      // 点击点位
      const index = mapState.points.findIndex(point => point.name === params.name);
      if (index !== -1) {
        mapState.points[index].selected = !mapState.points[index].selected;
        console.log('选中点位:', params.name, '状态:', mapState.points[index].selected);
      }
    }
  }
};

// 定义图表选项
const chartOption = computed(() => {
  return {
    backgroundColor: '#F3F3F3',
    tooltip: {
      trigger: 'item',
      formatter: (params: any) => {
        if (params.seriesType === 'map') {
          return `${params.name}: ${params.value || '暂无数据'}`;
        } else {
          return `${params.name}<br/>数值：${params.value[2]}`;
        }
      }
    },
    visualMap: {
      min: 0,
      max: 100,
      text: ['高', '低'],
      realtime: false,
      calculable: true,
      inRange: {
        color: ['#e0ffff', '#006edd']
      },
      textStyle: {
        color: '#000'
      }
    },
    geo: {
      map: 'customMap',
      roam: true, // 允许缩放和平移
      zoom: 1.5,
      label: {
        show: true,
        fontSize: 10,
        color: '#333'
      },
      itemStyle: {
        areaColor: '#CCCCCC',
        borderColor: '#666666',
        borderWidth: 1
      },
      emphasis: {
        label: {
          color: '#000'
        },
        itemStyle: {
          areaColor: '#F5DEB3'
        }
      }
    },
    series: [
      {
        name: '区域',
        type: 'map',
        map: 'customMap',
        selectedMode: 'single',
        geoIndex: 0,
        data: [
          { name: '区域A', value: 89 },
          { name: '区域B', value: 58 },
          { name: '区域C', value: 72 }
        ],
        select: {
          itemStyle: {
            areaColor: '#a6c84c'
          }
        }
      },
      {
        name: '点位',
        type: 'effectScatter',
        coordinateSystem: 'geo',
        symbolSize: (val: number[]) => val[2] * 1.5,
        showEffectOn: 'emphasis',
        rippleEffect: {
          brushType: 'stroke'
        },
        label: {
          show: true,
          formatter: '{b}',
          position: 'right',
          color: '#333'
        },
        itemStyle: {
          color: '#FF5722',
          shadowBlur: 5,
          shadowColor: '#333'
        },
        emphasis: {
          scale: true
        },
        data: mapState.points.map(point => ({
          name: point.name,
          value: point.value,
          itemStyle: point.selected ? { color: '#5C3317' } : undefined
        }))
      }
    ]
  };
});
</script>

<style scoped>
.map-container {
  width: 100%;
  height: 100%;
}

.chart {
  width: 100%;
  height: 600px;
}
</style>
