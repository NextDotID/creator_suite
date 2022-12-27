<template>
  <v-card width="100%">
    <v-list lines="two">
      <v-list-subheader inset> Folders </v-list-subheader>
      <v-list-item
        v-for="folder in dirs"
        :key="folder.name"
        :title="folder.path"
      >
        <template v-slot:prepend>
          <v-avatar color="grey-lighten-1" class="file-avatar">
            <!-- <v-icon color="white">mdi-folder</v-icon> -->
            <!-- <v-icon color="white">mdi-file-lock-open-outline</v-icon> -->
            <v-icon color="white" icon="mdi-folder-outline" />
          </v-avatar>
        </template>

        <v-list-item
          v-for="file in folder.children"
          :key="file.name"
          :title="file.name"
          :subtitle="file.path + ',' + file.size + ',' + file.update_time"
        >
          <template v-slot:prepend>
            <v-avatar>
              <v-icon color="white">mdi-lock-outline</v-icon>
            </v-avatar>
          </template>
        </v-list-item>

        <template v-slot:append>
          <v-btn
            color="grey-lighten-1"
            icon="mdi-information"
            variant="text"
          ></v-btn>
        </template>
      </v-list-item>

      <v-divider inset></v-divider>

      <v-list-subheader inset>Files</v-list-subheader>

      <v-list-item
        v-for="file in contents"
        :key="file.name"
        :title="file.name"
        :subtitle="file.path + ',' + file.size + ',' + file.update_time"
        lines="three"
      >
        <!-- :subtitle="file.path + file.size + file.update_time" -->
        <template v-slot:prepend>
          <v-avatar>
            <v-icon color="white">mdi-file-lock-open-outline</v-icon>
          </v-avatar>
        </template>
        <!-- <v-spacer></v-spacer> -->

        <template v-slot:append>
          <v-dialog v-model="dialog" persistent>
            <template v-slot:activator="{ props }">
              <v-btn
                color="primary"
                v-bind="props"
                @click="dialog = true"
                style="margin-right: 10px"
              >
                AES
                <v-icon end icon="mdi-lock-outline"></v-icon>
              </v-btn>
            </template>
            <encrypt
              class="dialog-encrypt"
              type="aes"
              :name="file.name"
              :origin_file="file.path"
            ></encrypt>
            <v-btn
              class="ma-2"
              color="green-darken-1"
              variant="text"
              @click="dialog = false"
            >
              Close
            </v-btn>
          </v-dialog>

          <v-btn class="ma-2" color="red">
            ECC
            <v-icon end icon="mdi-lock-outline"></v-icon>
          </v-btn>
          <v-btn class="ma-2" size="40" icon="mdi-share-variant"></v-btn>
          <v-btn class="ma-2" size="40" icon="mdi-rename-outline"></v-btn>
          <v-btn class="ma-2" size="40" icon="mdi-delete-outline"></v-btn>
          <v-btn color="grey-lighten-1" icon="mdi-information" variant="text">
            <!-- <v-icon>mdi-information</v-icon> -->
            <!-- <v-tooltip activator="parent" location="top"> TODO </v-tooltip> -->
            <v-menu>
              <template v-slot:activator="{ props }">
                <v-btn icon="mdi-information" v-bind="props"></v-btn>
              </template>
              <v-list>
                <v-list-item v-for="(a, i) in actions" :key="i">
                  <span>
                    <v-icon>{{ a.icon }}</v-icon>
                    <v-list-item-title>{{ a.title }}</v-list-item-title>
                  </span>
                </v-list-item>
              </v-list>
            </v-menu>
          </v-btn>
        </template>
      </v-list-item>
    </v-list>
  </v-card>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import encrypt from './Encrypt.vue'
// import FileService from '@/stores/file'
// import File from '@/types/File'
// import { Item, File, ListResponse } from '@/types/file'
import { storeToRefs } from 'pinia'
import { useBrowserStore } from '@/stores'

export default defineComponent({
  components: {
    encrypt,
  },
  props: {
    refreshPending: Boolean,
  },
  setup(props, ctx) {
    // 该入参包含了当前组件定义的所有 props
    console.log(props)
    const browser = useBrowserStore()
    // const { path } = storeToRefs(browser)
    const { folders } = storeToRefs(browser)
    const { files } = storeToRefs(browser)

    const load = async () => {
      browser.loadPath()
    }

    const refreshed = async () => {
      if (props.refreshPending) {
        await load()
        ctx.emit('refreshed', true)
      }
    }
    return { refreshed, folders, files }
  },
  data() {
    return {
      items: [],
      filter: '',
      dialog: false,
      dialog2: false,
      actions: [
        { title: 'Rename', icon: 'mdi-rename-outline' },
        { title: 'Share', icon: 'mdi-share-variant' },
        { title: 'Delete', icon: 'mdi-delete-outline' },
      ],
    }
  },
  computed: {
    dirs() {
      return this.folders
    },
    contents() {
      return this.files
    },
  },
  methods: {
    clickMe() {
      console.log('clicked')
      console.log(this.addNum(4, 2))
    },

    addNum(num1: number, num2: number): number {
      return num1 + num2
    },

    list() {},
  },
})
</script>

<style lang="less" scoped>
.dialog-encrypt {
  margin-left: 10%;
  width: 70%;
}
.file-avatar {
  background-image: url('@img/explore.png');
  background-position: 0px 0px;
  background-size: cover;
}
.menu {
  width: 500px;
}
</style>
