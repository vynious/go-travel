// AddConnectionModal.js
import React from 'react';

function AddConnectionModal({ isOpen, toggleModal }) {
    if (!isOpen) return null;

    return (
        <div className="fixed inset-0 z-50 overflow-auto bg-black bg-opacity-40 flex">
            <div className="relative p-8 bg-white w-full max-w-md m-auto flex-col flex rounded-lg">
                <div>
                    <h2 className="text-lg font-semibold">Find Users</h2>
                    <button onClick={toggleModal} className="absolute top-0 right-0 mt-4 mr-4">Close</button>
                </div>
                <input
                    type="text"
                    className="mt-4 p-2 border rounded-lg w-full"
                    placeholder="Search by name or email"
                />
                {/* Render search results and handle adding a connection here */}
            </div>
        </div>
    );
}

export default AddConnectionModal;
