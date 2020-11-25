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
    // @ts-ignore
    const res = validateTodoText(value)
    console.log(res)
    if (res && res !== '') {
      return false
    }
    return true
  },
  message() {
    return 'TODOは4文字以上、10文字以内です。'
  }
})
extend('required', required)
