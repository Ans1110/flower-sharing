import { LucideIcon } from "lucide-react";
import { Card, CardDescription, CardHeader, CardTitle } from "./card";
import { cn } from "@/lib/utils";

type CustomCardProps = {
  title: string;
  description: string;
  icon: LucideIcon;
  iconClassName?: string;
};

const CustomCard = ({
  title,
  description,
  icon: Icon,
  iconClassName,
}: CustomCardProps) => {
  return (
    <Card className="border-0 shadow-lg hover:shadow-xl transition-all duration-300 bg-white/80 dark:bg-black/80 backdrop-blur-sm">
      <CardHeader className="text-center">
        <div className="w-16 h-16 bg-linear-to-br from-rose-100 to-pink-100 dark:from-rose-100/20 dark:to-pink-100/30 rounded-full flex items-center justify-center mx-auto mb-4">
          <Icon className={cn("size-8", iconClassName)} />
        </div>
        <CardTitle className="text-2xl text-gray-900 dark:text-gray-100">
          {title}
        </CardTitle>
        <CardDescription className="text-gray-600 text-lg dark:text-gray-400">
          {description}
        </CardDescription>
      </CardHeader>
    </Card>
  );
};

export default CustomCard;
export { CustomCard };
