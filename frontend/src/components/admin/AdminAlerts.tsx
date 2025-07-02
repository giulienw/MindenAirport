import React from 'react';
import { AlertTriangle, Info, XCircle, Check, Clock } from 'lucide-react';
import type { SystemAlert } from '@/types';
import { adminService } from '@/services';

interface AdminAlertsProps {
  alerts: SystemAlert[];
  onResolveAlert: () => void;
  showAll?: boolean;
}

export const AdminAlerts: React.FC<AdminAlertsProps> = ({
  alerts,
  onResolveAlert,
  showAll = false,
}) => {
  const handleResolveAlert = async (alertId: string) => {
    try {
      await adminService.resolveAlert(alertId);
      onResolveAlert();
    } catch (error) {
      console.error('Failed to resolve alert:', error);
    }
  };

  const getAlertIcon = (type: SystemAlert['type']) => {
    switch (type) {
      case 'ERROR':
        return XCircle;
      case 'WARNING':
        return AlertTriangle;
      case 'INFO':
        return Info;
      default:
        return Info;
    }
  };

  const getAlertColor = (type: SystemAlert['type']) => {
    switch (type) {
      case 'ERROR':
        return 'text-red-600 bg-red-50 border-red-200';
      case 'WARNING':
        return 'text-orange-600 bg-orange-50 border-orange-200';
      case 'INFO':
        return 'text-blue-600 bg-blue-50 border-blue-200';
      default:
        return 'text-gray-600 bg-gray-50 border-gray-200';
    }
  };

  const getCategoryColor = (category: SystemAlert['category']) => {
    switch (category) {
      case 'FLIGHT':
        return 'bg-blue-100 text-blue-800';
      case 'BAGGAGE':
        return 'bg-purple-100 text-purple-800';
      case 'SYSTEM':
        return 'bg-gray-100 text-gray-800';
      case 'SECURITY':
        return 'bg-red-100 text-red-800';
      default:
        return 'bg-gray-100 text-gray-800';
    }
  };

  const formatTimeAgo = (timestamp: string) => {
    const now = new Date();
    const alertTime = new Date(timestamp);
    const diffInMinutes = Math.floor((now.getTime() - alertTime.getTime()) / (1000 * 60));

    if (diffInMinutes < 1) return 'Just now';
    if (diffInMinutes < 60) return `${diffInMinutes}m ago`;
    if (diffInMinutes < 1440) return `${Math.floor(diffInMinutes / 60)}h ago`;
    return `${Math.floor(diffInMinutes / 1440)}d ago`;
  };

  const activeAlerts = alerts.filter(alert => !alert.resolved);
  const resolvedAlerts = alerts.filter(alert => alert.resolved);

  return (
    <div className="bg-white rounded-lg shadow-sm border border-gray-200">
      <div className="p-6 border-b border-gray-200">
        <div className="flex items-center justify-between">
          <h3 className="text-lg font-semibold text-gray-900">System Alerts</h3>
          {!showAll && activeAlerts.length > 0 && (
            <span className="bg-red-100 text-red-800 text-xs font-medium px-2 py-1 rounded-full">
              {activeAlerts.length} active
            </span>
          )}
        </div>
      </div>

      <div className="p-6">
        {alerts.length === 0 ? (
          <div className="text-center py-8">
            <Check className="w-12 h-12 text-green-500 mx-auto mb-3" />
            <p className="text-gray-600">No alerts at this time</p>
            <p className="text-sm text-gray-500">System is running smoothly</p>
          </div>
        ) : (
          <div className="space-y-4">
            {showAll && (
              <>
                {activeAlerts.length > 0 && (
                  <div>
                    <h4 className="text-sm font-medium text-gray-900 mb-3">Active Alerts</h4>
                    <div className="space-y-3">
                      {activeAlerts.map((alert) => {
                        const Icon = getAlertIcon(alert.type);
                        return (
                          <div
                            key={alert.id}
                            className={`p-4 rounded-lg border ${getAlertColor(alert.type)}`}
                          >
                            <div className="flex items-start justify-between">
                              <div className="flex items-start space-x-3">
                                <Icon className="w-5 h-5 mt-0.5" />
                                <div className="flex-1">
                                  <div className="flex items-center space-x-2 mb-1">
                                    <span className={`text-xs font-medium px-2 py-1 rounded ${getCategoryColor(alert.category)}`}>
                                      {alert.category}
                                    </span>
                                    <span className="text-xs text-gray-500 flex items-center">
                                      <Clock className="w-3 h-3 mr-1" />
                                      {formatTimeAgo(alert.timestamp)}
                                    </span>
                                  </div>
                                  <p className="text-sm font-medium text-gray-900">
                                    {alert.message}
                                  </p>
                                </div>
                              </div>
                              <button
                                onClick={() => handleResolveAlert(alert.id)}
                                className="ml-4 text-xs px-3 py-1 bg-white border border-gray-300 rounded-md hover:bg-gray-50 transition-colors"
                              >
                                Resolve
                              </button>
                            </div>
                          </div>
                        );
                      })}
                    </div>
                  </div>
                )}

                {resolvedAlerts.length > 0 && (
                  <div>
                    <h4 className="text-sm font-medium text-gray-900 mb-3">Recently Resolved</h4>
                    <div className="space-y-3">
                      {resolvedAlerts.map((alert) => {
                        const Icon = getAlertIcon(alert.type);
                        return (
                          <div
                            key={alert.id}
                            className="p-4 rounded-lg border border-gray-200 bg-gray-50"
                          >
                            <div className="flex items-start space-x-3">
                              <Icon className="w-5 h-5 mt-0.5 text-gray-400" />
                              <div className="flex-1">
                                <div className="flex items-center space-x-2 mb-1">
                                  <span className="text-xs font-medium px-2 py-1 rounded bg-green-100 text-green-800">
                                    RESOLVED
                                  </span>
                                  <span className={`text-xs font-medium px-2 py-1 rounded ${getCategoryColor(alert.category)}`}>
                                    {alert.category}
                                  </span>
                                  <span className="text-xs text-gray-500 flex items-center">
                                    <Clock className="w-3 h-3 mr-1" />
                                    {formatTimeAgo(alert.timestamp)}
                                  </span>
                                </div>
                                <p className="text-sm text-gray-600 line-through">
                                  {alert.message}
                                </p>
                              </div>
                            </div>
                          </div>
                        );
                      })}
                    </div>
                  </div>
                )}
              </>
            )}

            {!showAll && (
              <div className="space-y-3">
                {alerts.slice(0, 5).map((alert) => {
                  const Icon = getAlertIcon(alert.type);
                  return (
                    <div
                      key={alert.id}
                      className={`p-3 rounded-lg border ${
                        alert.resolved
                          ? 'border-gray-200 bg-gray-50'
                          : getAlertColor(alert.type)
                      }`}
                    >
                      <div className="flex items-start justify-between">
                        <div className="flex items-start space-x-3">
                          <Icon className={`w-4 h-4 mt-0.5 ${alert.resolved ? 'text-gray-400' : ''}`} />
                          <div className="flex-1">
                            <div className="flex items-center space-x-2 mb-1">
                              {alert.resolved && (
                                <span className="text-xs font-medium px-2 py-1 rounded bg-green-100 text-green-800">
                                  RESOLVED
                                </span>
                              )}
                              <span className={`text-xs font-medium px-2 py-1 rounded ${getCategoryColor(alert.category)}`}>
                                {alert.category}
                              </span>
                              <span className="text-xs text-gray-500">
                                {formatTimeAgo(alert.timestamp)}
                              </span>
                            </div>
                            <p className={`text-sm ${alert.resolved ? 'text-gray-600 line-through' : 'text-gray-900 font-medium'}`}>
                              {alert.message}
                            </p>
                          </div>
                        </div>
                        {!alert.resolved && (
                          <button
                            onClick={() => handleResolveAlert(alert.id)}
                            className="ml-4 text-xs px-2 py-1 bg-white border border-gray-300 rounded hover:bg-gray-50 transition-colors"
                          >
                            Resolve
                          </button>
                        )}
                      </div>
                    </div>
                  );
                })}
              </div>
            )}
          </div>
        )}
      </div>
    </div>
  );
};
