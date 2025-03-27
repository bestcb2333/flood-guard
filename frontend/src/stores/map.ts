import apiAxios from "@/utils/axios";
import type {GeoJSON} from "echarts/types/src/coord/geo/geoTypes.js";
import { defineStore } from "pinia";

const useMapStore = defineStore('map', {
  state: (): {
    map: GeoJSON | null,
  } => ({
    map: null,
  }),
  actions: {
    async load(event?: string) {
      try {
        let url = event ? `/map?event=${event}` : '/map'
        const res = await apiAxios.get<GeoJSON>(url)
        this.map = res.data
      } catch (err) {
        console.log(err)
      }
    }
  },
})

export default useMapStore
