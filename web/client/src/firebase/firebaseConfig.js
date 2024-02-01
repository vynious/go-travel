
import { initializeApp, getApps, getApp } from 'firebase/app';
import 'firebase/auth';

// TODO: Add SDKs for Firebase products that you want to use


const firebaseConfig = {
    apiKey: "",
    authDomain: "",
    // the rest of firebase config
};


let firebaseApp;
if (!getApps().length) { 
    firebaseApp = initializeApp(firebaseConfig);
} else {
    firebaseApp = getApp();
}
