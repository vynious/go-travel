
import { initializeApp, getApps, getApp } from 'firebase/app';
import 'firebase/auth';


// TODO: Add SDKs for Firebase products that you want to use


const firebaseConfig = {
    apiKey: process.env.GATSBY_FB_API_KEY,
    authDomain: process.env.GATSBY_FB_AUTH_DOMAIN,
    projectId: process.env.GATSBY_FB_PROJECT_ID,
    storageBucket: process.env.GATSBY_FB_STORAGE_BUCKET,
    messagingSenderId: process.env.GATSBY_FB_MESSAGING_SENDER_ID,
    appId: process.env.GATSBY_FB_APP_ID,
};



let firebaseApp;
if (!getApps().length) {
    firebaseApp = initializeApp(firebaseConfig);
} 

export default firebaseApp;
