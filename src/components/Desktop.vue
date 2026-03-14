<script setup lang="ts">
import { ref, type Ref } from 'vue';
import HelloWorldApp from './apps/HelloWorldApp.vue';
import type { WindowInstance } from '../types';
import { Applications, DEFAULT_WINDOW_SIZE } from '../types';
import { generateId, getDefaultWindowPosition } from '../lib';

const windows: Ref<WindowInstance[]> = ref([]);

const removeWindow = (id: string) => {
  const index = windows.value.findIndex((w) => w.id === id);

  windows.value.splice(index, 1);
};

const addHelloWorldWindow = () => {
  windows.value.push({
    id: generateId(),
    type: Applications.HelloWorld,
    position: getDefaultWindowPosition(DEFAULT_WINDOW_SIZE),
    size: DEFAULT_WINDOW_SIZE,
  });
};
</script>

<template>
  <div class="absolute -z-10 inset-0 h-full w-full bg-[url(/background.jpg)] bg-cover bg-center" />

  <nav class="flex flex-col gap-2 p-4">
    <button
      class="p-2 rounded-lg bg-ef-aqua shadow-md w-fit h-fit cursor-pointer"
      @click="addHelloWorldWindow"
    >
      <p class="text-2xl">👋</p>
    </button>
  </nav>

  <div v-for="window in windows" :key="window.id">
    <HelloWorldApp
      v-if="window.type === Applications.HelloWorld"
      v-model:position="window.position"
      v-model:size="window.size"
      @closed="() => removeWindow(window.id)"
      :id="window.id"
    />
  </div>
</template>
