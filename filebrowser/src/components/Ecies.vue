<template>
  <v-card class="mx-auto" max-width="500">
    <v-card-title class="text-h6 font-weight-regular justify-space-between">
      <v-avatar color="primary" size="24" style="margin-right: 10px">
        {{ step }}
      </v-avatar>
      <span>{{ currentTitle }}</span>
    </v-card-title>

    <v-window v-model="step">
      <v-window-item :value="1">
        <v-card-text>
          <v-text-field prepend-icon="mdi-paperclip" label="Your File" placeholder="Your File"
            v-model="origin_file"></v-text-field>
          <span class="text-caption text-grey-darken-1">
            <br />
          </span>
          <v-file-input :rules="rules" accept="image/png, image/jpeg, image/bmp" placeholder="Pick an avatar"
            prepend-icon="mdi-camera" label="Avatar" v-model="selectedFiles">

          </v-file-input>
          <v-textarea color="#cb87f4" label="Decription" v-model="description"
            hint="Please add some words to share and promote"></v-textarea>
          <p :style="classObj">{{ message }}</p>
        </v-card-text>
      </v-window-item>

      <v-window-item :value="2">
        <v-card-text>
          <v-select :items="['mumbai']" label="Network" v-model="network"></v-select>
          <v-text-field label="Managed Contract" v-model="managed_contract"></v-text-field>
          <v-text-field label="Payment Token Address" v-model="payment_token_address"></v-text-field>
          <v-text-field label="Payment Token Amount" v-model="payment_token_amount"></v-text-field>
          <p :style="classObj">{{ message }}</p>
        </v-card-text>
      </v-window-item>
    </v-window>
    <v-divider></v-divider>
    <v-card-actions>
      <v-btn v-if="step > 1" variant="text" @click="back"> Back </v-btn>
      <v-spacer></v-spacer>
      <v-btn v-if="step == 1" class="vbutton" @click="save"> Save </v-btn>
      <v-btn v-if="step == 2" class="vbutton" @click="create_content"> Create </v-btn>
      <v-btn v-if="step < 2" @click="next">
        Next
        <v-icon>mdi-chevron-right</v-icon>
      </v-btn>
    </v-card-actions>
  </v-card>

</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { useBrowserStore } from '@/stores'
import http from '@/plugins/http-common'
import { CopyRequest } from '@/types/file'
import { CreateContentRequest, CreateContentResponse } from '@/types/content'
import ResponseData from '@/types/ResponseData'

export default defineComponent({
  props: {
    type: Number,
    name: String,
    extension: String,
    origin_file: String,
  },
  setup(props) {
    var extension = props.extension
    var origin_file = props.origin_file

    var rules: any = (value: string | any[]) => {
      return !value || !value.length || value[0].size < 2000000 || 'Avatar size should be less than 2 MB!'
    }

    const browser = useBrowserStore()
    return {
      browser,
      rules,
      origin_file,
      extension,
    }
  },
  data: () => ({
    create_msg: '',
    step: 1,
    message: '',
    message_color: 1,
    content_url: '',
    content_id: -1,
    managed_contract: '',
    payment_token_address: '',
    payment_token_amount: '',
    description: '',
    network: '',
    selectedFiles: [],
  }),
  computed: {
    classObj() {
      switch (this.message_color) {
        case 1:
          return { color: '#28C11C' }
        case 2:
          return { color: '#C11C1C' }
        case 3:
          return { color: '#D52CDE' }
        default:
          return { color: '#28C11C' }
      }
    },
    currentTitle() {
      switch (this.step) {
        case 1:
          return 'Choose file for Ecies encryption'
        case 2:
          return 'Create Content'
        default:
          return 'CreatorSuite'
      }
    },
  },
  methods: {
    back() {
      this.step--
      this.message = ''
    },
    next() {
      this.step++
      this.message = ''
    },
    save() {
      this.message = 'Save Content Succeeded'
      this.message_color = 1
    },
    create_content() {
      var req: CreateContentRequest = {
        content_locate_url: '',
        content_name: this.name!,
        managed_contract: this.managed_contract,
        network: this.network,
        payment_token_address: this.payment_token_address,
        payment_token_amount: this.payment_token_amount,
        key_id: -1,
        encryption_type: Number(this.type), // ecies
        file_extension: this.extension!,
        description: this.description
      }
      console.log(req)
      http
        .post('/api/v1/create', req)
        .then((response: CreateContentResponse) => {
          console.log(response.data)
          this.content_id = response.data.content_id
          this.message = 'CreateContent Succeeded'
          this.message_color = 1
          // move enc files to the storage/{id}
          var target =
            this.browser.path +
            '/' +
            this.content_id.toString() +
            '/' +
            this.name
          console.log('target: ', target)
          var move_req: CopyRequest = {
            src: this.origin_file!,
            dst: target,
          }
          http
            .post('api/v1/file/copy', move_req)
            .then((response: ResponseData) => {
              console.log(response.data)
            })
            .catch((e: Error) => {
              console.log(e)
            })
        })
        .catch((e: any) => {
          console.log(e)
          this.message = `CreateContent Failed: ' + "${e.response.data.message} "`
          this.message_color = 2
        })
    },
  }
})
</script>

<style lang="less" scoped>
.vbutton {
  // background-color: #29c11c;
  background-image: url('@img/explore.png');
  background-position: 0px 0px;
  background-size: cover;
  color: #000;
}
</style>
