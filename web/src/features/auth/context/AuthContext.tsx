import { createContext, useContext, useEffect, useState } from 'react';
import * as session from 'supertokens-web-js/recipe/session';

interface AuthContextType {
  isAuthenticated: boolean;
  user: any | null;
  loading: boolean;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [state, setState] = useState<AuthContextType>({
    isAuthenticated: false,
    user: null,
    loading: true,
  });

  useEffect(() => {
    // Initialize auth state
  }, []);

  return (
    <AuthContext.Provider value={state}>
      {children}
    </AuthContext.Provider>
  );
}

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
}; 