<script setup lang="ts">
import { ref, type Ref } from 'vue';
import HelloWorldApp from './apps/HelloWorldApp.vue'

const HelloWorld = "helloWorld";

type WindowType = typeof HelloWorld;

const windows: Ref<WindowType[]> = ref([]);

const removeWindow = (index: number) => windows.value.splice(index, 1);

const addHelloWorldWindow = () => {
  windows.value.push(HelloWorld);
}

</script>

<template>
  <div
    class="absolute -z-10 inset-0 h-full w-full bg-base-300 bg-[linear-gradient(to_right,#73737320_1px,transparent_1px),linear-gradient(to_bottom,#73737320_1px,transparent_1px)] bg-size-[20px_20px]"
  />

  <nav class="flex flex-col gap-2 p-4">
    <button class="p-2 rounded-lg bg-blue-200 shadow-md w-fit h-fit cursor-pointer" @click="addHelloWorldWindow">
      <p class="text-2xl">👋</p>
    </button>
  </nav>

  <div v-for="(window, index) in windows">
    <HelloWorldApp v-if="window === HelloWorld" @closed="() => removeWindow(index)"/>
  </div>

</template>
