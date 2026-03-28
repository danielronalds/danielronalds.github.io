<script setup lang="ts">
import Window from '@/components/Window.vue';
import type { Position, Size } from '@/types';
import { Hand } from 'lucide-vue-next';

defineProps<{ id: string }>();

defineEmits<{ closed: [] }>();

const isMinimized = defineModel('isMinimized', { default: false });
const position = defineModel<Position>('position', { required: true });
const size = defineModel<Size>('size', { required: true });
</script>

<template>
  <Window
    name="Hello world"
    v-if="!isMinimized"
    v-model:position="position"
    v-model:size="size"
    @closed="$emit('closed')"
    @minimized="isMinimized = true"
  >
    <template #icon>
      <Hand :size="16" />
    </template>

    <div class="w-full flex flex-col gap-6 justify-center items-center text-center p-10">
      <img
        src="/assets/avatar.jpg"
        class="size-36 rounded-full object-cover object-bottom-left mx-auto"
      />
      <p>👋 Hi, I'm Daniel. I like building stuff</p>
      <p>
        This is my personal website, poke around and have a look what's on it. The inspiration is
        for it to be a desktop like experience!
      </p>
    </div>
  </Window>
</template>
