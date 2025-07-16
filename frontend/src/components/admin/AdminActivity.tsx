import React from 'react';
import { Clock, User, Plane, Package, Settings, Shield } from 'lucide-react';
import type { AdminActivity as AdminActivityType } from '@/types';

interface AdminActivityProps {
  activities: AdminActivityType[];
}

export const AdminActivity: React.FC<AdminActivityProps> = ({ activities }) => {
  const getCategoryIcon = (category: AdminActivityType['category']) => {
    switch (category) {
      case 'USER':
        return User;
      case 'FLIGHT':
        return Plane;
      case 'BAGGAGE':
        return Package;
      case 'SYSTEM':
        return Settings;
      default:
        return Shield;
    }
  };

  const getCategoryColor = (category: AdminActivityType['category']) => {
    switch (category) {
      case 'USER':
        return 'bg-blue-100 text-blue-800';
      case 'FLIGHT':
        return 'bg-green-100 text-green-800';
      case 'BAGGAGE':
        return 'bg-purple-100 text-purple-800';
      case 'SYSTEM':
        return 'bg-gray-100 text-gray-800';
      default:
        return 'bg-orange-100 text-orange-800';
    }
  };

  const formatTimeAgo = (timestamp: string) => {
    const now = new Date();
    const activityTime = new Date(timestamp);
    const diffInMinutes = Math.floor((now.getTime() - activityTime.getTime()) / (1000 * 60));

    if (diffInMinutes < 1) return 'Just now';
    if (diffInMinutes < 60) return `${diffInMinutes}m ago`;
    if (diffInMinutes < 1440) return `${Math.floor(diffInMinutes / 60)}h ago`;
    return `${Math.floor(diffInMinutes / 1440)}d ago`;
  };

  return (
    <div className="bg-white rounded-lg shadow-sm border border-gray-200">
      <div className="p-6 border-b border-gray-200">
        <div className="flex items-center justify-between">
          <h3 className="text-lg font-semibold text-gray-900">Recent Activity</h3>
          <Clock className="w-5 h-5 text-gray-400" />
        </div>
      </div>

      <div className="p-6">
        {activities.length === 0 ? (
          <div className="text-center py-8">
            <Clock className="w-12 h-12 text-gray-400 mx-auto mb-3" />
            <p className="text-gray-600">No recent activity</p>
            <p className="text-sm text-gray-500">Admin actions will appear here</p>
          </div>
        ) : (
          <div className="space-y-4">
            {activities.map((activity) => {
              const Icon = getCategoryIcon(activity.category);
              return (
                <div key={activity.id} className="flex items-start space-x-3 p-3 rounded-lg hover:bg-gray-50 transition-colors">
                  <div className={`p-2 rounded-lg ${getCategoryColor(activity.category)}`}>
                    <Icon className="w-4 h-4" />
                  </div>
                  <div className="flex-1 min-w-0">
                    <div className="flex items-center justify-between mb-1">
                      <p className="text-sm font-medium text-gray-900">
                        {activity.action}
                      </p>
                      <span className="text-xs text-gray-500">
                        {formatTimeAgo(activity.timestamp)}
                      </span>
                    </div>
                    <p className="text-sm text-gray-600 mb-1">
                      {activity.details}
                    </p>
                    <div className="flex items-center space-x-2">
                      <span className={`text-xs font-medium px-2 py-1 rounded ${getCategoryColor(activity.category)}`}>
                        {activity.category}
                      </span>
                      <span className="text-xs text-gray-500">
                        by {activity.adminName}
                      </span>
                    </div>
                  </div>
                </div>
              );
            })}
          </div>
        )}
      </div>

      {activities.length > 0 && (
        <div className="px-6 py-3 border-t border-gray-200">
          <button className="text-sm text-blue-600 hover:text-blue-700 font-medium">
            View all activity
          </button>
        </div>
      )}
    </div>
  );
};
