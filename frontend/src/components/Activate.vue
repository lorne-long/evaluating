<script lang="ts" setup>
import {reactive} from 'vue'
import {Greet1} from '../../wailsjs/go/main/App'
const data = reactive({
  name: "",
  resultText: "Please enter your Activate Code",
})

function greet() {
  Greet1(data.name).then(result => {
    data.resultText = result
  })
}

</script>

<template>
  <main>
    <div id="result" class="result">{{ data.resultText }}</div>
    <div id="input" class="input-box">
      <input id="name" v-model="data.name" autocomplete="off" class="input" type="text"/>
      <button class="btn" @click="greet">Greet</button>
    </div>
  </main>
</template>

<style scoped>
</style>
