<template>
  <v-container>
    <ValidationObserver
      ref="observer"
      v-slot="{ invalid }"
      @submit.prevent="submit"
      slim
    >
      <v-form>
        <v-row>
          <v-col cols="5">
            <v-card>
              <ValidationProvider
                v-slot="{ errors }"
                name="TODO"
                rules="todoText"
              >
                <v-text-field
                  v-model="todoInput.text"
                  :error-messages="errors[0]"
                  label="TODO"></v-text-field>
              </ValidationProvider>
            </v-card>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="5">
            <v-card>
              <ValidationProvider
                v-slot="{ errors }"
                name="ユーザーID"
                rules="required"
              >
                <v-text-field
                  v-model="todoInput.userId"
                  :error-messages="errors[0]"
                  label="ユーザーID"></v-text-field>
              </ValidationProvider>
            </v-card>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="5">
            <v-btn @click="submit">登録</v-btn>
          </v-col>
        </v-row>
      </v-form>
    </ValidationObserver>
  </v-container>
</template>
<script lang="ts">
import { Component, Vue } from '~/node_modules/nuxt-property-decorator'
import { ValidationObserver } from 'vee-validate'
import { NewTodo } from '~/gql-types'

export class TodoInput implements NewTodo {
  text: string = ''
  userId: string = ''
}

@Component({})
export default class TodoForm extends Vue {
  $refs!: {
    observer: InstanceType<typeof ValidationObserver>
  }

  todoInput: NewTodo = new TodoInput()

  async submit() {
    console.log('Submit!')
    const valid = await this.$refs.observer.validate()
    if (valid) {
      console.log('valid!!!')
      // MEMO: GraphQLを介してバックエンドに登録内容を送信。（今回の検証範囲外なので省略）
    }
  }
}
</script>
