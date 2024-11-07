import { ErrorBoundary } from '@/shared/components/ErrorBoundary';
import { AuthProvider } from '@/features/auth/context/AuthContext';

export function Providers({ children }: { children: React.ReactNode }) {
  return (
    <ErrorBoundary>
      <AuthProvider>
        {children}
      </AuthProvider>
    </ErrorBoundary>
  );
} 