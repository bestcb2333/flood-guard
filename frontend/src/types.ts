export interface Model {
  id: number,
  createdAt: Date,
  updatedAt: Date,
}

export interface User extends Model {
  name: string,
  email: string,
  profile: string | null,
  admin: boolean,
  region: Region | null,
  notices: Notice[],
  events: Event[],
  analyses: Analysis[],
}

export interface Region extends Model {
  name: string,
  description: string | null,
  altitude: number | null,
  drainage: number | null,
  forecast: number | null,
  coordinate: MultiPolygon,
  users: User[],
  events: Event[],
  histories: History[],
  resources: Resource[],
  sensors: Sensor[],
}

export interface Event extends Model {
  regions: Region[],
  startTime: Date,
  endTime: Date,
  user: User | null,
  severity: number,
  coordinate: MultiPolygon,
  description: string | null,
}

export interface History {
  id: number,
  createdAt: Date,
  region: Region,
  rainfall: number | null,
  waterlevel: number | null,
  description: string | null,
}

export interface Resource extends Model {
  type: string,
  name: string,
  quantity: number,
  region: Region,
  coordinate: Point,
  available: boolean,
}

export interface Notice extends Model {
  user: User | null,
  title: string,
  content: string | null,
}

export interface Sensor extends Model {
  name: string,
  coordinate: Point,
  description: string | null,
  available: boolean,
  region: Region,
}

export interface Analysis extends Model {
  id: number,
  createdAt: Date,
  user: User,
  role: string,
  content: string | null,
}

export interface FieldMeta {
  comment: string,
  type: string,
}

export interface TableData<T> {
  count: number,
  fields: Record<string, FieldMeta>,
  data: T[],
}

export type Point = [x: number, y: number]
export type Ring = Point[]
export type MultiPoint = Point[]
export type Polygon = Ring[]
export type MultiPolygon = Polygon[]

export interface MultiPolygonGeometry {
  type: "MultiPolygon",
  coordinates: MultiPolygon,
}

export interface Feature {
  type: "Feature",
  properties: Record<string, any>,
  geometry: MultiPolygonGeometry,
}

export interface FeatureCollection {
  type: 'FeatureCollection',
  features: Feature[],
}
