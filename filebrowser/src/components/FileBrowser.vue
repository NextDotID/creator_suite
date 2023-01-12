<template>
  <!-- <span> {{ path }} </span> -->
  <div class="toolbar">
    <Toolbar></Toolbar>
  </div>
  <div class="filelist">
    <List :refreshed="true"></List>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { useBrowserStore } from '@/stores'
import Toolbar from './Toolbar.vue'
import List from './List.vue'

export default defineComponent({
  setup() {
    const browser = useBrowserStore()
    const load = async () => {
      browser.loadPath()
    }
    return {
      load
    }
  },
  async mounted() {
    await this.load()
  }
})

// import { storeToRefs } from 'pinia'

// const browser = useBrowserStore()
// const { path } = storeToRefs(browser)




</script>

<!-- <script lang="ts">
import { defineComponent } from 'vue'
import Toolbar from './Toolbar.vue'
import List from './List.vue'

export default defineComponent({
  data: () => ({
    step: 1,
    path: '/storage',
  }),
  components: {
    Toolbar,
    List,
  },
  // setup() {
  // },
  computed: {
    currentTitle() {
      switch (this.step) {
        case 1:
          return 'Sign-up'
        case 2:
          return 'Create a password'
        default:
          return 'Account created'
      }
    },
  },
  methods: {
    onChangePath(v: string) {
      this.path = v
    },
  },
})
</script> -->

<style lang="less" scoped>
.toolbar {
  display: flex;
  background-color: white;
}

.filelist {
  display: flex;
  width: 100%;
  height: 100%;
  background-color: white;
}
</style>
