<script setup lang="ts">
import { computed } from 'vue';
import { useDraggable } from '../composables/useDraggable';
import { useResizable } from '../composables/useResizable';
import type { Position, Size } from '../types';
import { DEFAULT_WINDOW_SIZE } from '../types';
import { getDefaultWindowPosition } from '../lib';

import { Minus, X } from 'lucide-vue-next';

const { name = 'Window' } = defineProps<{ name?: string }>();

defineEmits<{ closed: []; minimized: [] }>();

const position = defineModel<Position>('position', {
  default: () => getDefaultWindowPosition(DEFAULT_WINDOW_SIZE),
});
const size = defineModel<Size>('size', { default: () => ({ ...DEFAULT_WINDOW_SIZE }) });

const { onMouseDown: onDragMouseDown } = useDraggable(position);
const { onMouseDown: onResizeMouseDown, isResizing } = useResizable(size);

const style = computed(() => ({
  top: `${position.value.y}px`,
  left: `${position.value.x}px`,
  width: `${size.value.width}px`,
  height: `${size.value.height}px`,
}));
</script>

<template>
  <div class="bg-cp-bg0 absolute shadow-xl overflow-y-hidden rounded-xl" :style="style">
    <header
      @mousedown="onDragMouseDown"
      class="p-2 bg-cp-bg1 cursor-grab active:cursor-grabbing select-none flex flex-row justify-between"
    >
      <div class="flex flex-row w-fit gap-2 items-center">
        <slot name="icon" />
        <p class="font-semibold">{{ name }}</p>
      </div>

      <div class="flex flex-row gap-2">
        <button @click="$emit('minimized')" class="cursor-pointer">
          <minus :size="18" />
        </button>
        <button @click="$emit('closed')" class="cursor-pointer">
          <X :size="18" />
        </button>
      </div>
    </header>

    <div class="overflow-y-auto" :class="{ 'select-none': isResizing }">
      <slot />
    </div>

    <div
      @mousedown="onResizeMouseDown"
      class="absolute bottom-0 right-0 w-3 h-3 cursor-nwse-resize"
    />
  </div>
</template>
