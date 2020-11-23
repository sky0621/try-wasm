<template>
  <TodoForm />
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import TodoForm from '~/components/TodoForm.vue'
@Component({
  components: { TodoForm },
})
export default class IndexPage extends Vue {}
</script>
