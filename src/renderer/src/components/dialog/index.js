import { createApp, h } from 'vue'
import AppDialog from "./AppDialog.vue"

let instance = null

const Modal = {
  show(options) {
    if (instance) {
      return instance.show(options);
    }
    const container = document.createElement('div');
    document.body.appendChild(container);

    const app = createApp({
      render() {
        return h(AppDialog, {
          ref: 'dialog'
        })
      }
    })
    instance = app.mount(container).$refs.dialog;
    return instance.show(options);
  }
}
export {
  AppDialog
}
export default Modal