import { create, StateCreator } from "zustand";
import { immer } from "zustand/middleware/immer";
import { persist, createJSONStorage } from "zustand/middleware";

type StoreConfig<T> = {
  name?: string;
  storage?: Storage;
  skipPersist?: boolean;
  excludeFromPersist?: Array<keyof T>;
};

/**
 * Create a Zustand store implementing immer middleware with SSR safe storage, persisting only on the client.
 * @param initializer - The initializer function
 * @param config - The configuration object
 * @returns The Zustand store
 */
export default function createStore<T extends object>(
  initializer: StateCreator<T, [["zustand/immer", never]], [], T>,
  config: StoreConfig<T> = {}
) {
  const { name, storage, skipPersist, excludeFromPersist } = config;

  const immerWapper = immer(initializer);

  const isClient = typeof globalThis.window !== "undefined";
  const safeStorage = isClient ? storage : undefined;

  if (skipPersist || !safeStorage) {
    return create(immerWapper);
  }

  return create(
    persist(immerWapper, {
      name: name ?? "zustand-store",
      storage: createJSONStorage(() => safeStorage),
      partialize: (state) => {
        return Object.fromEntries(
          Object.entries(state).filter(
            ([key]) => !excludeFromPersist?.includes(key as keyof T)
          )
        );
      },
    })
  );
}

export { createStore };
