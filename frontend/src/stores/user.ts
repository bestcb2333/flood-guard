import type {User} from "@/types";
import {defineStore} from "pinia";

export const useUserStore = defineStore('user', {
  state: (): {val: User | null} => ({
    val: null,
  })
})

export default useUserStore
