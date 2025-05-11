export interface User {
  id: number
  createdAt: Date
  updatedAt: Date
  name: string
  email: string
  avatar: string
  profile: string
  admin: boolean
  region: Region | null
  notices: Notice[]
  events: Event[]
  analyses: Analysis[]
}

export interface Region {
  id: number
  updatedAt: Date
  name: string
  description: string
  altitude: number | null
  drainage: number | null
  forecast: number | null
  coordinate: [number, number][][][]
  users: User[]
  events: Event[]
  histories: History[]
  resources: Resource[]
  sensors: Sensor[]
}

export interface Event {
  id: number
  createdAt: Date
  name: string
  regions: Region[]
  startTime: Date
  endTime: Date
  user: User | null
  severity: number
  coordinate: [number, number]
  description: string
}

export interface History {
  id: number
  createdAt: Date
  region: Region
  rainfall: number | null
  waterlevel: number | null
  description: string
}

export interface Resource {
  id: number
  updatedAt: Date
  type: string
  name: string
  quantity: number
  region: Region
  coordinate: [number, number]
  available: boolean
}

export interface Wiki {
  id: number
  name: string
  notices: Notice[]
}

export interface Notice {
  id: number
  createdAt: Date
  updatedAt: Date
  user: User | null
  title: string
  content: string
}

export interface Sensor {
  id: number
  createdAt: Date
  name: string
  coordinate: [number, number]
  description: string
  available: boolean
  region: Region
}

export interface Analysis {
  id: number
  createdAt: Date
  user: User
  role: string
  content: string
}
