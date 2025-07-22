import { Copy } from "lucide-react";
import { useURLHook } from "../provider/url.provider";
import { CopyToClipboard } from "../utils/utils";
export default function Display() {
  const { url } = useURLHook();
  return (
    <div className="flex flex-col gap-2 my-5 d items-center mx-auto">
      <h2 className="text-xl  font-medium">Your URL will be visible here</h2>
      <div className="flex gap-10 justify-between items-center mx-auto border border-gray-300 min-w-full p-2">
        <a
          href={url?.endpoint}
          className="text-sm lg:text-base"
          target="_blank"
        >
          {url?.endpoint}
        </a>
        <Copy
          className="min-w-5 min-h-5 text-gray-500"
          onClick={() => CopyToClipboard(url?.endpoint || "")}
        ></Copy>
      </div>
    </div>
  );
}
