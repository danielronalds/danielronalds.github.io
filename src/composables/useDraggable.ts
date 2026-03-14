import { ref, type Ref } from 'vue';
import type { Position } from '../types';

export const useDraggable = (position: Ref<Position>) => {
  const isDragging = ref(false);
  const lastMousePosition = ref<Position>({ x: 0, y: 0 });

  const onMouseUp = () => {
    isDragging.value = false;
    document.removeEventListener('mousemove', onMouseMove);
    document.removeEventListener('mouseup', onMouseUp);
  };

  const onMouseMove = (event: MouseEvent) => {
    if (!isDragging.value) return;

    position.value = {
      x: position.value.x + event.clientX - lastMousePosition.value.x,
      y: position.value.y + event.clientY - lastMousePosition.value.y,
    };

    lastMousePosition.value = { x: event.clientX, y: event.clientY };
  };

  const onMouseDown = (event: MouseEvent) => {
    isDragging.value = true;
    lastMousePosition.value = { x: event.clientX, y: event.clientY };
    document.addEventListener('mousemove', onMouseMove);
    document.addEventListener('mouseup', onMouseUp);
  };

  return { onMouseDown };
};
