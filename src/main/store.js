import Store from 'electron-store'
const store = new Store()

export const getConfig = (key) => {
  return store.get(key)
}
export const setConfig = (key, value) => {
  return store.set(key, value)
}

export default store