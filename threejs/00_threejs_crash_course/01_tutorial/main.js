import * as THREE from 'three';
import { OrbitControls } from 'three/examples/jsm/Addons.js';

// initialize a scene
const scene = new THREE.Scene();

// add object to a scene

const cubeGeometry = new THREE.BoxGeometry(1,1,1);
const cubeMaterial = new THREE.MeshBasicMaterial({ color: "red"});
const cubeMesh = new THREE.Mesh(
  cubeGeometry,
  cubeMaterial
);
scene.add(cubeMesh);

// initialize the camera

const camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 30);
camera.position.z = 5;

// we can add a camera to a scene 
// scene.add(camera);



// Initialize the renderer
const canvas = document.querySelector('canvas.threejs');
const renderer = new THREE.WebGLRenderer({canvas});
renderer.setSize(window.innerWidth, window.innerHeight);

renderer.render(scene, camera);