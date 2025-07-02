import React from 'react';
import { AdminDashboard, AdminGuard } from '@/components/admin';

export const AdminPage: React.FC = () => {
  return (
    <AdminGuard>
      <AdminDashboard />
    </AdminGuard>
  );
};
