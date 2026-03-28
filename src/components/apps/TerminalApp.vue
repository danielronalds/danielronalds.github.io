<script setup lang="ts">
import Window from '@/components/Window.vue';
import type { Position, Size } from '@/types';
import { SquareTerminal } from 'lucide-vue-next';
import { useTerminal } from '@/composables/useTerminal';
import { onMounted, useTemplateRef } from 'vue';
import Prompt from './Terminal/Prompt.vue';

defineProps<{ id: string }>();

defineEmits<{ closed: [] }>();

const isMinimized = defineModel('isMinimized', { default: false });
const position = defineModel<Position>('position', { required: true });
const size = defineModel<Size>('size', { required: true });

const terminal = useTerminal();

const prompt = useTemplateRef('prompt');

onMounted(() => {
  prompt.value?.focus();
});
</script>

<template>
  <Window
    name="Terminal"
    v-if="!isMinimized"
    v-model:position="position"
    v-model:size="size"
    @closed="$emit('closed')"
    @minimized="isMinimized = true"
  >
    <template #icon>
      <SquareTerminal :size="16" />
    </template>

    <div class="w-full font-mono p-2">
      <div class="flex flex-row gap-0">
        <Prompt :path="terminal.path" />
        <input
          type="text"
          ref="prompt"
          class="w-full caret-black caret-block text-sm outline-none text-cp-fg p-0"
        />
      </div>
    </div>
  </Window>
</template>

<style scoped>
.caret-block {
  caret-shape: block;
}
</style>
