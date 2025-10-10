
<script lang="ts" setup>
import { ref, reactive, onMounted, onBeforeUnmount } from 'vue'
import { Greet } from '../../wailsjs/go/main/App'
import '../assets/styles/HelloWorld.css'
// --- Reactive data for greeting (unchanged)
const data = reactive({
  name: "",
  resultText: "Please enter your name below 👇",
})

function greet() {
  Greet(data.name).then(result => {
    data.resultText = result
  })
}
</script>

<template>
  <main>
    <!-- Greet Section -->
    <div id="result" class="result">{{ data.resultText }}</div>
    <div id="input" class="input-box">
      <input id="name" v-model="data.name" autocomplete="off" class="input" type="text"/>
      <button class="btn" @click="greet">Greet</button>
    </div>
  </main>
</template>
