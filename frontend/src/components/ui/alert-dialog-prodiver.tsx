import { useGlobalStore } from "@/hooks/useGlobalStore";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "./alert-dialog";

const AlertDialogProvider = () => {
  const { alertOpen, alertConfig, updateAlertOpen } = useGlobalStore();

  if (!alertConfig) return null;

  const handleConfirm = () => {
    if (alertConfig?.onConfirm) alertConfig.onConfirm();
    updateAlertOpen(false);
  };

  const handleCancel = () => {
    if (alertConfig?.onCancel) alertConfig.onCancel();
    updateAlertOpen(false);
  };

  return (
    <AlertDialog open={alertOpen} onOpenChange={updateAlertOpen}>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>
            {alertConfig.title || "Confirmation Required"}
          </AlertDialogTitle>
          <AlertDialogDescription>
            {alertConfig.description || "Are you sure you want to Action?"}
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel onAbort={handleCancel}>
            {alertConfig.cancelLabel || "Cancel"}
          </AlertDialogCancel>
          <AlertDialogAction onClick={handleConfirm}>
            {alertConfig.confirmLabel || "Continue"}
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
};

export default AlertDialogProvider;
export { AlertDialogProvider };
