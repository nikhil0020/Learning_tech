import * as THREE from "three";
import { OrbitControls } from "three/addons/controls/OrbitControls.js";
import { Pane } from "tweakpane";

// initialize the pane
const pane = new Pane();

// initialize the scene
const scene = new THREE.Scene();

const axesHelper = new THREE.AxesHelper(6);
scene.add(axesHelper);

// initialize the geometry
const geometry = new THREE.BoxGeometry(1, 1, 1);
const planeGeometry = new THREE.PlaneGeometry(1,1, 2);
const torusKnotGeometry = new THREE.TorusKnotGeometry(0.5, 0.15, 100, 16);
const sphereGeometry = new THREE.SphereGeometry(0.5, 32, 32);
const cylinderGeometry = new THREE.CylinderGeometry(0.5, 0.5, 1, 32);
// initialize the material
// const material = new THREE.MeshBasicMaterial({
//   color: 'limeGreen',
//   transparent: true,
//   opacity: 0.5
// });


// const material = new THREE.MeshLambertMaterial();
// const material = new THREE.MeshStandardMaterial(); // This is industry practice and used across unity, unreal , etc.
const material = new THREE.MeshPhysicalMaterial();

pane.addBinding(material, 'roughness', {
  min: 0,
  max: 1,
  step: 0.01
})

pane.addBinding(material, 'metalness', {
  min: 0,
  max: 1,
  step: 0.01
})

// const material = new THREE.MeshBasicMaterial();

// material.color = 'limeGreen'; // This will not work because color is a object
// material.color = new THREE.Color(0xff1234)
// material.transparent = true;
// material.opacity = 1;
material.side = THREE.DoubleSide;
// material.fog = true;

// const fog = new THREE.Fog(0x0fff234, 1, 10);
// scene.fog = fog;
// scene.background = new THREE.Color(0x0fff234)
// initialize the mesh
const mesh = new THREE.Mesh(geometry, material);
const mesh2 = new THREE.Mesh(torusKnotGeometry, material);
mesh2.position.x = 1.5;
const plane = new THREE.Mesh(planeGeometry, material);
plane.position.x = -1.5;
const sphere = new THREE.Mesh(sphereGeometry, material);
sphere.position.y = 1.5;
const cylinder = new THREE.Mesh(cylinderGeometry, material);
cylinder.position.y = -1.5;

const group = new THREE.Group();
group.add(mesh, mesh2, plane, sphere, cylinder);
// scene.add(mesh);
// scene.add(mesh2);
// scene.add(plane);
// scene.add(sphere);
// scene.add(cylinder);
// scene.add(mesh, mesh2, plane, sphere, cylinder);
scene.add(group);

// initialize the light
const light = new THREE.AmbientLight(0xffffff, 0.2);
scene.add(light);

const pointLight = new THREE.PointLight(0xffffff, 100);
pointLight.position.set(5,5,5)
scene.add(pointLight);

// initialize the camera
const camera = new THREE.PerspectiveCamera(
  35,
  window.innerWidth / window.innerHeight,
  0.1,
  200
);
camera.position.z = 5;

// initialize the renderer
const canvas = document.querySelector("canvas.threejs");
const renderer = new THREE.WebGLRenderer({
  canvas: canvas,
  antialias: true,
});
renderer.setSize(window.innerWidth, window.innerHeight);
renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2));

// instantiate the controls
const controls = new OrbitControls(camera, canvas);
controls.enableDamping = true;

window.addEventListener("resize", () => {
  camera.aspect = window.innerWidth / window.innerHeight;
  camera.updateProjectionMatrix();
  renderer.setSize(window.innerWidth, window.innerHeight);
});

// render the scene
const renderloop = () => {

  group.children.forEach((child) => {
    if ( child instanceof THREE.Mesh) {
      child.rotation.y += 0.01;
    }
  })

  controls.update();
  renderer.render(scene, camera);
  window.requestAnimationFrame(renderloop);
};

renderloop();
