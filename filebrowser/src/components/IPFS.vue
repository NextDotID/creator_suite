<template>
  <v-card class="mx-auto" max-width="500">
    <v-card-title class="text-h6 font-weight-regular justify-space-between">
      <span> IPFS Connection</span>
    </v-card-title>

    <v-card-text>
      <v-text-field
        label="Host"
        placeholder="http://localhost"
        v-model="host"
      ></v-text-field>
      <v-text-field
        label="API Port"
        placeholder="5001"
        v-model="api_port"
      ></v-text-field>
      <v-text-field
        label="GateWay Port"
        placeholder="8080"
        v-model="gateway_port"
      ></v-text-field>
      <v-text-field
        label="PeerID"
        type="password"
        v-model="peer_id"
      ></v-text-field>
      <v-text-field
        label="Public Key"
        type="password"
        v-model="pub_key"
      ></v-text-field>
      <span class="text-caption text-grey-darken-1">
        Please enter IPFS Connection Params
      </span>
      <span class="text-caption text-red-darken-1">
        {{ message }}
      </span>
    </v-card-text>
    <v-card-actions>
      <v-btn variant="text" @click="conntect"> Connect </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script setup lang="ts">
// import { ref } from 'vue'
import { storeToRefs } from 'pinia'
import { useIpfsStore } from '@/stores'
const store = useIpfsStore()
const { message } = storeToRefs(store)
const { peer_id } = storeToRefs(store)
const { pub_key } = storeToRefs(store)
const { host } = storeToRefs(store)
const { api_port } = storeToRefs(store)
const { gateway_port } = storeToRefs(store)

// defineProps<{ msg: string }>()

const conntect = () => {
  store.checkAliveSync(
    peer_id.value,
    pub_key.value,
    host.value,
    api_port.value,
    gateway_port.value
  )
}
</script>
