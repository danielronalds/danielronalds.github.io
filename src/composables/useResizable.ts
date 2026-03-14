import { ref, type Ref } from 'vue';
import type { Size } from '../types';

export const useResizable = (size: Ref<Size>, minWidth = 100, minHeight = 100) => {
  const isResizing = ref(false);
  const lastMousePosition = ref({ x: 0, y: 0 });

  const onMouseUp = () => {
    isResizing.value = false;
    document.removeEventListener('mousemove', onMouseMove);
    document.removeEventListener('mouseup', onMouseUp);
  };

  const onMouseMove = (event: MouseEvent) => {
    if (!isResizing.value) return;

    size.value = {
      width: Math.max(minWidth, size.value.width + event.clientX - lastMousePosition.value.x),
      height: Math.max(minHeight, size.value.height + event.clientY - lastMousePosition.value.y),
    };

    lastMousePosition.value = { x: event.clientX, y: event.clientY };
  };

  const onMouseDown = (event: MouseEvent) => {
    isResizing.value = true;
    lastMousePosition.value = { x: event.clientX, y: event.clientY };
    document.addEventListener('mousemove', onMouseMove);
    document.addEventListener('mouseup', onMouseUp);
  };

  return { onMouseDown, isResizing };
};
