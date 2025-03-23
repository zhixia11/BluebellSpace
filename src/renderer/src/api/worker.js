import http from "./http"

export const submit = (name, config) => {
  return http.post('/worker/submit', { name, config })
}

export const progressEventSource = (id) => {
  return new EventSource(import.meta.env.VITE_API_URL + '/worker/progress?id=' + id)
}

export const cancel = (id) => {
  return http.get('/worker/cancel', {
    params: { id }
  })
}