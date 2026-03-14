export type Position = { x: number; y: number };

export type Size = { width: number; height: number };

export const DEFAULT_WINDOW_SIZE: Size = { width: 500, height: 400 };

export interface WindowInstance {
  id: string;
  type: string;
  position: Position;
  size: Size;
}

export enum Application {
  HelloWorld = 'helloWorld',
}
