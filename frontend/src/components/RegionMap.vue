<script setup lang="ts">
import {MapChart} from 'echarts/charts';
import {TitleComponent, TooltipComponent, VisualMapComponent} from 'echarts/components';
import {use} from 'echarts/core';
import {CanvasRenderer} from 'echarts/renderers';
import type {GeoJSON} from 'echarts/types/src/coord/geo/geoTypes.js';
import {watch} from 'vue';
import {ref} from 'vue';
import * as echarts from 'echarts';
import VChart from 'vue-echarts';
import {onMounted} from 'vue';
import type {Point} from '@/types';

use([CanvasRenderer, MapChart, TooltipComponent, TitleComponent, VisualMapComponent])

const chartOptions = ref<echarts.EChartsOption>({})

export interface MapPoint {
  name: string,
  coordinate: Point,
}

export interface Data {
  name: string,
  value: number,
}

const props = defineProps<{
  name: string,
  geojson: GeoJSON | null,
  points?: MapPoint[],
  data: Data[],
  min?: number,
  max?: number,
}>()

onMounted(() => {
  watch(props, (props) => {
    if (props.geojson) {
      echarts.registerMap(props.name, props.geojson)
      chartOptions.value = {
        tooltip: { trigger: "item" },
        visualMap: {
          min: props.min ?? 0,
          max: props.max ?? 1,
          left: "left",
          top: "bottom",
          text: ["高", "低"],
          calculable: true,
          inRange: {
            color: ['lightskyblue', 'yellow', 'orangered'],
          },
        },
        series: [
          {
            type: "map",
            map: props.name,
            roam: true,
            label: { show: true },
            data: props.data,
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
            data: props.points?.map(point => ({
              name: point.name,
              value: point.coordinate,
              itemStyle: { color: '#5C3317' }
            }))
          }
        ],
      }
    }
  }, {immediate: true})
})

</script>

<template>
  {{ props.points }}
  <v-chart :option="chartOptions" autoresize />
</template>
