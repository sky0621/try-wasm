import Vue from 'vue'
import {
  extend,
  ValidationObserver,
  ValidationProvider
} from 'vee-validate'
import { required } from 'vee-validate/dist/rules'

Vue.component('ValidationProvider', ValidationProvider)
Vue.component('ValidationObserver', ValidationObserver)

extend('todoText', {
  validate: (value: string) => {
    if (value.length < 4) {
      return false
    }
    if (value.length > 10) {
      return false
    }
    return true
  },
  message() {
    return 'TODOは4文字以上、10文字以内です。'
  }
})
extend('required', required)
