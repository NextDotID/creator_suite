<template>
  <v-card class="toolbar-card">
    <v-toolbar class="toolbar">
      <v-btn variant="text" icon="mdi:mdi-menu"></v-btn>
      <!-- <v-toolbar-title> {{ path }}</v-toolbar-title> -->
      <v-text-field
        class="input-field"
        hide-details
        prepend-icon="mdi-file-search-outline"
        single-line
        v-model="path"
        @keydown.enter="refreshed"
        @click="refreshed"
      ></v-text-field>
      <v-btn variant="text" icon="mdi-magnify" v-on:click="refreshed"></v-btn>
      <!-- <v-btn class="ma-2" color="primary">
        Input File
        <v-icon end icon="mdi-checkbox-marked-circle"></v-icon>
      </v-btn> -->
      <!-- <v-btn class="ma-2" color="red">
        Decline
        <v-icon end icon="mdi-cancel"></v-icon>
      </v-btn>

      <v-btn class="ma-2">
        <v-icon start icon="mdi-minus-circle"></v-icon>
        Cancel
      </v-btn>

      <v-btn class="ma-2" color="orange-darken-2">
        <v-icon start icon="mdi-arrow-left"></v-icon>
        Back
      </v-btn>

      <v-btn class="ma-2" color="purple" icon="mdi-wrench"></v-btn>
      <v-btn class="ma-2" color="indigo" icon="mdi-cloud-upload"></v-btn> -->

      <!-- <v-btn
        class="ma-2"
        variant="text"
        icon="mdi-thumb-up"
        color="blue-lighten-2"
      ></v-btn>

      <v-btn
        class="ma-2"
        variant="text"
        icon="mdi-thumb-down"
        color="red-lighten-2"
      ></v-btn> -->

      <!-- <v-btn class="ma-2" variant="text" color="white-lighten-2">
        <v-icon>{{ icons.mdiAccount }}</v-icon>
      </v-btn>
      <v-btn class="ma-2" variant="text" color="white-lighten-2">
        <v-icon>{{ icons.mdiPencil }}</v-icon>
      </v-btn>
      <v-btn class="ma-2" variant="text" color="white-lighten-2">
        <v-icon>{{ icons.mdiShareVariant }}</v-icon>
      </v-btn>
      <v-btn class="ma-2" variant="text" color="white-lighten-2">
        <v-icon>{{ icons.mdiDelete }}</v-icon>
      </v-btn>
      <v-btn class="ma-2" variant="text" color="white-lighten-2">
        <v-icon>{{ icons.mdiFileLockOpenOutline }}</v-icon>
      </v-btn>
      <v-btn class="ma-2" variant="text" color="white-lighten-2">
        <v-icon>{{ icons.mdiFileLockOutline }}</v-icon>
      </v-btn> -->
      <!-- <v-btn variant="text" icon="mdi:mdi-menu"></v-btn>
      <v-toolbar-title>My files</v-toolbar-title> -->
      <v-spacer></v-spacer>
      <!-- <v-btn variant="text" icon="mdi-view-module"></v-btn> -->

      <v-dialog v-model="ipfsdialog">
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" @click="ipfsdialog = true" class="ipfs_button">
            <v-icon>mdi-lan-connect</v-icon>
            IPFS
          </v-btn>
        </template>
        <ipfs class="dialog-encrypt"></ipfs>
        <v-btn
          class="ma-2"
          color="green-darken-1"
          variant="text"
          @click="ipfsdialog = false"
        >
          Close
        </v-btn>
      </v-dialog>
    </v-toolbar>
  </v-card>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import ipfs from '@cp/IPFS.vue'
import { useBrowserStore } from '@/stores'
import { storeToRefs } from 'pinia'

export default defineComponent({
  components: {
    ipfs,
  },
  setup() {
    const browser = useBrowserStore()
    const { path } = storeToRefs(browser)
    // const load = async () => {
    //   browser.loadPath()
    // }
    // const refreshed = async () => {
    //   await load()
    // }

    const load = () => {
      browser.loadPath()
    }
    const refreshed = () => {
      console.log('try loading')
      load()
    }
    return { refreshed, path }
  },
  data() {
    return {
      ipfsdialog: false,
      icons: {
        mdiAccount: 'mdi-account',
        mdiPencil: 'mdi-pencil',
        mdiShareVariant: 'mdi-share-variant',
        mdiDelete: 'mdi-delete',
        mdiFileLockOutline: 'mdi-file-lock-outline',
        mdiFileLockOpenOutline: 'mdi-file-lock-open-outline',
      },
    }
  },
  methods: {
    tryThis() {
      console.log('trying')
    },
    clickMe() {
      console.log('clicked')
      console.log(this.addNum(4, 2))
    },

    addNum(num1: number, num2: number): number {
      return num1 + num2
    },
  },
})
</script>

<style lang="less" scoped>
.navbutton {
  background-image: url('@img/explore.png');
  background-position: 0px 0px;
  background-size: cover;
  margin-left: 10px;
}

// .ipfs_button {
//   background-color: whitesmoke;
// }

.dialog-encrypt {
  margin-left: 10%;
  width: 70%;
}

.toolbar-card {
  // width: 80%;
  // margin-left: 10%;
  width: 100%;
}
.toolbar {
  background-image: url('@img/explore.png');
  background-position: 0px 0px;
  background-size: cover;
}
</style>
