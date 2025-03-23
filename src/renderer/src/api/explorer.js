import http from "./http"

export const deleteFile = (paths) => {
  if (!Array.isArray(paths)) {
    paths = [paths]
  }
  return http.post('/explorer/delete', { paths })
}

export const info = (path) => {
  return http.get('/explorer/info', {
    params: { path }
  })
}