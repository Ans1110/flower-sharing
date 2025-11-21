import createStore from "@/lib/createStore";

type AlertConfigType = {
  title: string;
  description: string;
  confirmLabel: string;
  cancelLabel: string;
  onConfirm: () => void;
  onCancel: () => void;
};

type StateType = {
  alertOpen: boolean;
  alertConfig: AlertConfigType | null;
};

type ActionType = {
  updateAlertOpen: (isOpen: StateType["alertOpen"]) => void;
  showAlert: (config: AlertConfigType) => void;
};

type StoreType = StateType & ActionType;

const useGlobalStore = createStore<StoreType>(
  (set) => ({
    alertOpen: false,
    alertConfig: null,
    updateAlertOpen: (isOpen: StateType["alertOpen"]) =>
      set((state) => {
        state.alertOpen = isOpen;
        if (!isOpen) state.alertConfig = null;
      }),
    showAlert: (config: AlertConfigType) =>
      set((state) => {
        state.alertConfig = config;
        state.alertOpen = true;
      }),
  }),
  {
    name: "global-store",
    excludeFromPersist: ["alertOpen"],
  }
);

const alert = (config: AlertConfigType) => {
  useGlobalStore.getState().showAlert(config);
};

export { useGlobalStore, alert };
