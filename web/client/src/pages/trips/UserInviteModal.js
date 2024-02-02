import React, { useState } from 'react';

function UserInviteModal({ isOpen, onClose, onAddUser }) {
    if (!isOpen) return null;

    return (
        <div className="fixed inset-0 z-50 overflow-y-auto">
            <div className="flex items-end justify-center min-h-screen px-4 pt-4 pb-20 text-center sm:block sm:p-0">
                <div className="fixed inset-0 transition-opacity" onClick={onClose}>
                    <div className="absolute inset-0 bg-gray-500 opacity-75"></div>
                </div>
                {/* Modal content */}
                <span className="hidden sm:inline-block sm:align-middle sm:h-screen">&#8203;</span>
                <div className="inline-block overflow-hidden text-left align-bottom transition-all transform bg-white rounded-lg shadow-xl sm:my-8 sm:align-middle sm:max-w-lg sm:w-full">
                    {/* Modal body */}
                    <div className="px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
                        <h3 className="text-lg font-medium leading-6 text-gray-900" id="modal-title">Invite your team</h3>
                        {/* Form fields and buttons */}
                    </div>
                    <div className="px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse">
                        <button type="button" className="w-full px-4 py-2 mt-3 font-medium text-white bg-blue-600 rounded-md shadow-sm sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm hover:bg-blue-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500" onClick={onAddUser}>
                            Send invites
                        </button>
                        <button type="button" className="w-full px-4 py-2 mt-3 font-medium text-gray-700 bg-white rounded-md shadow-sm sm:mt-0 sm:w-auto sm:text-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500" onClick={onClose}>
                            Cancel
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default UserInviteModal;
