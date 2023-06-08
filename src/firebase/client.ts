import { initializeApp } from "firebase/app";

const firebaseConfig = {
  apiKey: "AIzaSyCJ6lK6GYh5yfO4YklsG-RGEAsx7KLdW5M",
  authDomain: "stat-n-track.firebaseapp.com",
  projectId: "stat-n-track",
  storageBucket: "stat-n-track.appspot.com",
  messagingSenderId: "644756878013",
  appId: "1:644756878013:web:6fa64a93ef50571fc8fdd4"
};

export const app = initializeApp(firebaseConfig);
