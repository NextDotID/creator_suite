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
          <!-- <v-text-field
            label="Email"
            placeholder="john@google.com"
          ></v-text-field> -->
          <v-text-field
            label="Encrypt File"
            placeholder="encrypt_file"
            v-model="encrypt_file"
          ></v-text-field>
          <v-text-field
            label="Password"
            type="password"
            v-model="key"
          ></v-text-field>
          <span class="text-caption text-grey-darken-1">
            Please enter a password
          </span>
          <p :style="classObj">{{ message }}</p>
        </v-card-text>
      </v-window-item>

      <v-window-item :value="2">
        <v-card-text>
          <!-- <v-text-field label="Password" type="password"></v-text-field>
          <v-text-field label="Confirm Password" type="password"></v-text-field> -->
          <v-text-field label="Host" v-model="store.host"></v-text-field>
          <v-text-field
            label="API Port"
            v-model="store.api_port"
          ></v-text-field>
          <v-text-field
            label="GateWay Port"
            v-model="store.gateway_port"
          ></v-text-field>
          <v-text-field
            label="PeerID"
            type="password"
            v-model="store.peer_id"
          ></v-text-field>
          <v-text-field
            label="Public Key"
            type="password"
            v-model="store.pub_key"
          ></v-text-field>
          <span class="text-caption text-grey-darken-1">
            Please enter IPFS Connection Params
          </span>
          <p :style="classObj">{{ message }}</p>
        </v-card-text>
      </v-window-item>

      <v-window-item :value="3">
        <div class="pa-4 text-center">
          <!-- <h3 class="text-h6 font-weight-light mb-2">Welcome to Vuetify</h3>
          <span class="text-caption text-grey">Thanks for signing up!</span> -->
          <v-text-field
            label="Location Url"
            v-model="content_url"
          ></v-text-field>
          <v-text-field label="Key ID" v-model="key_id"></v-text-field>
          <v-text-field
            label="Managed Contract"
            v-model="managed_contract"
          ></v-text-field>
          <v-text-field
            label="Payment Token Address"
            v-model="payment_token_address"
          ></v-text-field>
          <v-text-field
            label="Payment Token Amount"
            v-model="payment_token_amount"
          ></v-text-field>
          <p :style="classObj">{{ message }}</p>
        </div>
      </v-window-item>
    </v-window>

    <v-divider></v-divider>

    <v-card-actions>
      <v-btn v-if="step > 1" variant="text" @click="back"> Back </v-btn>
      <v-spacer></v-spacer>
      <v-btn v-if="step == 1" class="vbutton" @click="create"> Encrypt </v-btn>
      <v-btn v-if="step == 2" class="vbutton" @click="upload"> Upload </v-btn>
      <v-btn v-if="step == 3" class="vbutton" @click="create_content">
        Create
      </v-btn>
      <!-- <v-btn v-if="step < 3" color="primary" variant="flat" @click="step++">
        Next
      </v-btn> -->
      <v-btn v-if="step < 3" @click="next">
        Next
        <v-icon>mdi-chevron-right</v-icon>
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { useIpfsStore, useBrowserStore } from '@/stores'
import http from '@/plugins/http-common'
import { IpfsCfg, IpfsUpload, IpfsUploadResult } from '@/types/ipfs'
import { CreateRequest, CreateResponse, MoveRequest } from '@/types/file'
import { CreateContentRequest, CreateContentResponse } from '@/types/content'
import ResponseData from '@/types/ResponseData'

// const peer_id = storeToRefs(store)
// const pub_key = storeToRefs(store)
// const host = storeToRefs(store)
// const api_port = storeToRefs(store)
// const gateway_port = storeToRefs(store)
export default defineComponent({
  props: {
    type: String,
    name: String,
    origin_file: String,
  },
  setup(props) {
    var encrypt_file = props.origin_file + '.enc'
    var encrypt_name = props.name + '.enc'
    console.log(props,'dialog')
    const store = useIpfsStore()
    const browser = useBrowserStore()
    return {
      store,
      browser,
      encrypt_name,
      encrypt_file,
    }
  },
  data: () => ({
    create_msg: '',
    step: 1,
    key: '',
    key_id: -1,
    message: '',
    message_color: 1,
    content_url: '',
    content_id: -1,
    managed_contract: '',
    payment_token_address: '',
    payment_token_amount: 0,
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
          return 'Encrypt file by password'
        case 2:
          return 'Upload encrypted file to IPFS'
        case 3:
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
    create() {
      var req: CreateRequest = {
        encrypt_type: this.type!,
        key: this.key,
        origin_file: this.origin_file!,
        encrypt_file: this.encrypt_file,
      }
      console.log(req)
      http
        .post('/api/v1/file/create', req)
        .then((response: CreateResponse) => {
          console.log(response.data)
          console.log(response.data.encrypt_file)
          this.encrypt_file = response.data.encrypt_file
          this.key_id = response.data.key_id
          this.message = 'Encrypt content finished! key id is ' + this.key_id
          this.message_color = 1
        })
        .catch((e: Error) => {
          console.log(e)
          this.message = `Encrypt Failed ' + "${e.response.data.message} "`
          this.message_color = 2
        })
    },
    create_content() {
      var req: CreateContentRequest = {
        content_locate_url: this.content_url,
        managed_contract: this.managed_contract,
        payment_token_address: this.payment_token_address,
        payment_token_amount: Number(this.payment_token_amount),
        key_id: Number(this.key_id),
      }
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
            this.encrypt_name
          console.log('target: ', target)
          var move_req: MoveRequest = {
            src: this.encrypt_file,
            dst: target,
          }
          http
            .post('api/v1/file/move', move_req)
            .then((response: ResponseData) => {
              console.log(response.data)
            })
            .catch((e: Error) => {
              console.log(e)
            })
        })
        .catch((e: Error) => {
          console.log(e)
          this.message = `CreateContent Failed: ' + "${e.response.data.message} "`
          this.message_color = 2
        })
    },
    upload() {
      var cfg: IpfsCfg = this.store.getIpfsCfg()
      var req: IpfsUpload = {
        local_file: this.encrypt_file,
        cfg: cfg,
      }
      console.log(req)
      http
        .post('/api/v1/ipfs/upload', req)
        .then((response: IpfsUploadResult) => {
          console.log(response.data)
          console.log(response.data.content_locate_url)
          this.content_url = response.data.content_locate_url
          this.message = 'Upload Succeeded: ' + response.data.content_locate_url
          this.message_color = 1
        })
        .catch((e: Error) => {
          console.log(e)
          this.content_url = ''
          this.message = `Upload Failed: ' + "${e.response.data.message} "`
          this.message_color = 2
        })
    },
  },
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
