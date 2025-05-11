import dayjs from 'dayjs'

export function formatDate(_1: any, _2: any, cellValue: string): string {
  return dayjs(cellValue).format('MM月DD日 HH:mm')
}
