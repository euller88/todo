<template>
  <v-container fluid>
    <v-fab-transition>
      <v-btn
        color="success"
        class="v-btn--example"
        absolute
        fab
        dark
        large
        bottom
        right
        @click.stop="dialog = true"
      >
        <v-icon>mdi-plus</v-icon>
      </v-btn>
    </v-fab-transition>
    <v-data-iterator
      :items="items"
      item-key="created_at"
      :items-per-page.sync="itemsPerPage"
      :page.sync="page"
      hide-default-footer
    >
      <template #default="props">
        <v-row>
          <v-col v-for="item in props.items" :key="item.created_at" cols="12">
            <v-card>
              <div class="d-flex flex-no-wrap justify-space-between">
                <div>
                  <v-card-title primary-title>
                    {{ item.content }}
                  </v-card-title>
                  <v-card-subtitle>
                    Created at {{ getDate(item.created_at) }}
                  </v-card-subtitle>
                </div>
                <v-avatar class="ma-3" size="100" tile>
                  <v-btn
                    class="ma-3"
                    large
                    text
                    icon
                    color="red"
                    @click.stop="deleteTodo(item.id)"
                  >
                    <v-icon large>mdi-delete</v-icon>
                  </v-btn>
                </v-avatar>
              </div>
            </v-card>
          </v-col>
        </v-row>
      </template>
      <template #footer>
        <v-row class="mt-2" align="center" justify="center">
          <span class="grey--text">Items per page</span>
          <v-menu offset-y>
            <template #activator="{ on, attrs }">
              <v-btn
                dark
                text
                color="primary"
                class="ml-2"
                v-bind="attrs"
                v-on="on"
              >
                {{ itemsPerPage }}
                <v-icon>mdi-chevron-down</v-icon>
              </v-btn>
            </template>
            <v-list>
              <v-list-item
                v-for="(number, index) in itemsPerPageArray"
                :key="index"
                @click="updateItemsPerPage(number)"
              >
                <v-list-item-title>{{ number }}</v-list-item-title>
              </v-list-item>
            </v-list>
          </v-menu>

          <v-spacer></v-spacer>

          <span class="mr-4 grey--text">
            Page {{ page }} of {{ numberOfPages }}
          </span>
          <v-btn
            fab
            dark
            color="blue darken-3"
            class="mr-1"
            @click="formerPage"
          >
            <v-icon>mdi-chevron-left</v-icon>
          </v-btn>
          <v-btn fab dark color="blue darken-3" class="ml-1" @click="nextPage">
            <v-icon>mdi-chevron-right</v-icon>
          </v-btn>
        </v-row>
      </template>
    </v-data-iterator>
    <v-dialog
      v-model="dialog"
      :overlay="false"
      max-width="500px"
      transition="dialog-transition"
    >
      <v-card>
        <v-card-title primary-title>New To Do</v-card-title>
        <v-card-text>
          <v-text-field
            v-model="todo"
            label="What is needed to do?"
          ></v-text-field>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="error" @click="dialog = false">Cancel</v-btn>
          <v-btn color="success" @click="createTodo">Create</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
export default {
  name: 'IndexPage',
  async asyncData({ $axios }) {
    let items = []
    try {
      items = await $axios.$get('/todos')
    } catch (error) {
      console.log(error)
      items = []
    }
    return { items }
  },
  data() {
    return {
      items: [],
      itemsPerPageArray: [4, 8, 12],
      page: 1,
      itemsPerPage: 4,
      dialog: false,
      todo: '',
    }
  },
  computed: {
    numberOfPages() {
      return Math.ceil(this.items.length / this.itemsPerPage)
    },
  },
  methods: {
    async listTodos() {
      try {
        const items = await this.$axios.$get('/todos')
        this.items = items
      } catch (error) {
        console.log(error)
        this.items = []
      }
    },
    async deleteTodo(id) {
      try {
        await this.$axios.$delete(`/todos/${id}`)
      } catch (error) {
        console.log(error)
      } finally {
        await this.listTodos()
      }
    },
    async createTodo() {
      const payload = { content: this.todo }
      try {
        await this.$axios({
          url: '/todos',
          method: 'POST',
          data: payload,
        })
      } catch (error) {
        console.log(error)
      } finally {
        this.todo = ''
        this.dialog = false
        await this.listTodos()
      }
    },
    getDate(num) {
      const d = new Date(0)
      d.setUTCSeconds(num)
      return d.toLocaleString()
    },
    nextPage() {
      if (this.page + 1 <= this.numberOfPages) this.page += 1
    },
    formerPage() {
      if (this.page - 1 >= 1) this.page -= 1
    },
    updateItemsPerPage(number) {
      this.itemsPerPage = number
    },
  },
}
</script>

<style>
#lateral .v-btn--example {
  bottom: 0;
  position: absolute;
  margin: 0 0 16px 16px;
}
</style>