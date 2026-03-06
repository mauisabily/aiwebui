import axios from 'axios'

export function useApi() {
  const api = axios.create({ baseURL: '/api/v1' })
  return api
}