<template>
  <div class="container">
    <v-hover v-slot="{ isHovering, props }">
      <v-card class="mx-auto" max-width="400" v-bind="props">
        <v-img src="https://cdn.vuetifyjs.com/images/cards/forest-art.jpg"></v-img>

        <v-card-text>
          <h2 class="text-h6 text-primary">
            {{ file_name || 'Content Name' }}
          </h2>
          {{ file_description || 'default file description' }}
        </v-card-text>

        <v-card-title>
          <v-rating :model-value="4" dense color="orange" background-color="orange" hover class="mr-2"></v-rating>
          <span class="text-primary text-subtitle-2">size: {{ file_size || '1kb' }}</span>
        </v-card-title>

        <v-overlay :model-value="isHovering" contained scrim="#036358" class="align-center justify-center">
          <v-btn variant="flat">Purchase <v-icon color="blue">mdi-ethereum</v-icon>{{ payment_token_amount }}
            for more info</v-btn>
        </v-overlay>
      </v-card>
    </v-hover>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs } from 'vue'
import http from '@/plugins/http-common'
import { ShowContentRequest, ShowContentResponse } from '@/types/content'
export default defineComponent({
  // props: {
  //   content_id: Number
  // },
  setup(props) {
    const data = reactive({
      file_name: 'Content Name',
      file_description: 'default file description',
      file_size: '64KB',
      network: 'unknown',
      payment_token_amount: '1',
      payment_token_address: 'Token Address',
    })
    // var preview_image: Buffer = Buffer.alloc(0)
    const refData = toRefs(data)
    return {
      ...refData,
    }
  },
  mounted() {
    console.log('mounted!')
    var req: ShowContentRequest = {
      content_id: 38
    }
    http
      .get('/api/v1/show-content?content_id=38')
      .then((response: ShowContentResponse) => {
        console.log(response.data)
        this.file_name = response.data.content_name
        // this.preview_image = response.data.preview_image
        this.file_description = response.data.description
        this.file_size = response.data.file_size

        this.network = response.data.network
        this.payment_token_amount = response.data.payment_token_amount
        this.payment_token_address = response.data.payment_token_address
      })
  },
  data: () => ({
    overlay: false,
  }),
  methods: {
    load() {

    }
  }

})
// export default {
//   data: () => ({
//     overlay: false,
//   }),
// }

</script>

<style>
.container {
  height: 100vh;
  display: flex;
  align-items: center;
}
</style>
