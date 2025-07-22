import { useContext, createContext, type ReactNode, useState } from "react";
import { type URLType } from "../app.types";

interface URLContextType {
  url: URLType | null;
  setURL: (url: URLType | null) => void;
}

const URLContext = createContext<URLContextType | null>(null);

export function URLProvider({ children }: { children: ReactNode }) {
  const [url, setURL] = useState<URLType | null>(null);
  return (
    <URLContext.Provider value={{ url, setURL }}>
      {children}
    </URLContext.Provider>
  );
}

export const useURLHook = () => {
  const context = useContext(URLContext);
  if (!context) throw new Error("useURL can't be used outside URLProvider");
  return context;
};
