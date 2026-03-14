<script setup lang="ts">
import { ref, type Ref } from 'vue'
import HelloWorldApp from './apps/HelloWorldApp.vue'
import type { WindowInstance } from '../types'
import { DEFAULT_WINDOW_POSITION, DEFAULT_WINDOW_SIZE } from '../types'

const HelloWorld = 'helloWorld'

const windows: Ref<WindowInstance[]> = ref([])

const removeWindow = (id: string) => {
  const index = windows.value.findIndex((w) => w.id === id)

  windows.value.splice(index, 1)
}

const generateWindowId = () => {
  const alphabet = 'abcdefghijklmnopqrstuvwxyz'
  let randomString = ''

  for (let i = 0; i < 10; i++) {
    randomString += alphabet[Math.floor(Math.random() * alphabet.length)]
  }

  return randomString
}

const addHelloWorldWindow = () => {
  windows.value.push({
    id: generateWindowId(),
    type: HelloWorld,
    position: DEFAULT_WINDOW_POSITION,
    size: DEFAULT_WINDOW_SIZE,
  })
}
</script>

<template>
  <div
    class="absolute -z-10 inset-0 h-full w-full bg-base-300 bg-[linear-gradient(to_right,#73737320_1px,transparent_1px),linear-gradient(to_bottom,#73737320_1px,transparent_1px)] bg-size-[20px_20px]"
  />

  <nav class="flex flex-col gap-2 p-4">
    <button
      class="p-2 rounded-lg bg-blue-200 shadow-md w-fit h-fit cursor-pointer"
      @click="addHelloWorldWindow"
    >
      <p class="text-2xl">👋</p>
    </button>
  </nav>

  <div v-for="window in windows" :key="window.id">
    <HelloWorldApp
      v-if="window.type === HelloWorld"
      v-model:position="window.position"
      v-model:size="window.size"
      @closed="() => removeWindow(window.id)"
      :id="window.id"
    />
  </div>
</template>
