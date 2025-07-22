import { useCallback, useMemo, useState } from "react";
import useURL from "../hooks/useURL";
import { useURLHook } from "../provider/url.provider";

function MethodCard({
  onClick,
  label,
  active,
}: {
  onClick: Function;
  label: string;
  active: string[];
}) {
  const isActive = useMemo(() => {
    return active.includes(label);
  }, [active, label]);

  return (
    <div
      onClick={() => onClick(label)}
      className={`p-1 rounded-md flex items-center justify-center cursor-pointer ${
        isActive ? "bg-gray-300" : "bg-gray-100"
      }`}
    >
      {label}
    </div>
  );
}

export default function Form() {
  const [text, setText] = useState<string>("");
  const [timestamp, setTimeStamp] = useState<string>("");
  const [method, setMethod] = useState<string[]>([]);
  const { CreateURL } = useURL();
  const { setURL } = useURLHook();

  const updateMethod = useCallback((label: string) => {
    setMethod((prev) =>
      prev.includes(label) ? prev.filter((m) => m !== label) : [...prev, label]
    );
  }, []);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const res = await CreateURL({
      url: text,
      expireIn: timestamp,
      method: method,
    });
    if (!res.success) {
      alert(res.message);
      return;
    }
    setURL({ endpoint: res.endpoint });
  };

  return (
    <div className="flex w-full">
      <form onSubmit={handleSubmit} className="w-full flex flex-col gap-4">
        <div className="flex flex-col md:flex-row gap-4">
          <input
            type="text"
            value={text}
            onChange={(e) => setText(e.target.value)}
            placeholder="https://....."
            className="w-full py-2 px-4 border border-gray-400 rounded-md focus:outline-none focus:ring-0 focus:border-none"
          />

          <input
            type="datetime-local"
            value={timestamp}
            onChange={(e) => setTimeStamp(e.target.value)}
            className="py-2 px-4 border border-gray-400 rounded-md focus:outline-none focus:ring-0 focus:border-none"
          />
        </div>

        <div className="flex gap-2 flex-wrap">
          {["GET", "OPTION", "DELETE", "POST", "PATCH", "PUT"].map((ele) => (
            <MethodCard
              key={ele}
              active={method}
              label={ele}
              onClick={updateMethod}
            />
          ))}
        </div>
        <button className="bg-blue-400 py-2 px-4 text-white">Submit</button>
      </form>
    </div>
  );
}
