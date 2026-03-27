<script setup lang="ts">
import { ref, type Ref } from 'vue';
import HelloWorldApp from './apps/HelloWorldApp.vue';
import type { WindowInstance } from '../types';
import { Application, DEFAULT_WINDOW_SIZE } from '../types';
import { generateId, getDefaultWindowPosition } from '../lib';
import HelloWorldIcon from './apps/HelloWorldIcon.vue';

const windows: Ref<WindowInstance[]> = ref([]);

const removeWindow = (id: string) => {
  const index = windows.value.findIndex((w) => w.id === id);

  windows.value.splice(index, 1);
};

const launchApplication = (application: Application) => {
  windows.value.push({
    id: generateId(),
    type: application,
    position: getDefaultWindowPosition(DEFAULT_WINDOW_SIZE),
    size: DEFAULT_WINDOW_SIZE,
  });
};

// Starting the hello world one by default
launchApplication(Application.HelloWorld);
</script>

<template>
  <div
    class="absolute -z-10 inset-0 h-full w-full bg-cover bg-center"
  />

  <nav class="flex flex-col gap-2 p-4">
    <HelloWorldIcon @click="() => launchApplication(Application.HelloWorld)" />
  </nav>

  <div v-for="window in windows" :key="window.id">
    <HelloWorldApp
      v-if="window.type === Application.HelloWorld"
      v-model:position="window.position"
      v-model:size="window.size"
      @closed="() => removeWindow(window.id)"
      :id="window.id"
    />
  </div>
</template>
