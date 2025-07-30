// use std::sync::{Arc, Mutex};
// use std::thread;

// fn main() {
//     let counter = Arc::new(Mutex::new(0));

//     let counter_for_thread = Arc::clone(&counter);

//     let handle = thread::spawn(move || {
//         let mut num = counter_for_thread.lock().unwrap();
//         *num += 1;
//         println!("  Thread: counter = {}", *num);
//     });

//     handle.join().unwrap();
//     {
//         let mut num = counter.lock().unwrap();
//         *num += 10;
//         println!("  main thread: counter = {}", *num);
//     }


      
//     println!(" the end : {}", *counter.lock().unwrap());
// }


// use std::thread;
// use std::time::Duration;

// fn say_hello() {
//     for i in 1..=5 {
//         println!("hi  it's i :{}", i);
//         thread::sleep(Duration::from_millis(1000));
//     }
// }

// fn main() {
//     let handle = thread::spawn(|| {
//         say_hello();
//     });

//     println!(" start Thread...");

//     handle.join().unwrap();
//     println!(" Thread end ...");
    
// }

use raster::{Image, Color};
use std::f64::consts::PI;
use std::{thread, time::Duration};

fn main() {
    let mut image = Image::blank(500, 500);
    let red = Color::rgb(255, 0, 0);

    draw_heart(&mut image, 250.0, 250.0, 10.0, red.clone()); 

}

fn draw_heart(image: &mut Image, center_x: f64, center_y: f64, scale: f64, color: Color) {
    let segments = 700;
    let step = 2.0 * PI / segments as f64;

    for i in 0..=segments {
        
        thread::sleep(Duration::from_millis(50));
        
        
        let t = i as f64 * step;
        
        let x = 16.0 * (t.sin()).powi(3);
        
        let y = 13.0 * t.cos() - 5.0 * (2.0 * t).cos() - 2.0 * (3.0 * t).cos() - (4.0 * t).cos();
        
        let px = center_x + x * scale;
        let py = center_y - y * scale; 
        
        let _ = image.set_pixel(px.round() as i32, py.round() as i32, color.clone());
        raster::save(&image, "heart.png").expect("Failed to save image");
    }
}

// use raster::{Image,Color};
// use std::f64::consts::PI;
// use std::{thread, time::Duration};
// fn main() {

//     let mut image = Image::blank(500, 500);
//     let red = Color::rgb(255, 255, 255);

//     draw_circle(&mut image, 150, 150 , 80, red.clone());
// }





// fn draw_circle(image: &mut Image, center_x: i32, center_y: i32, radius: i32, color: Color) {
//     let segments = 360;

    
//     let c: f64 = 2.0 * PI / segments as f64;  

//     for i in 1..=segments {

//         thread::sleep(Duration::from_millis(80));

//         let t = c * i as f64 ;
//         let x = center_x as f64 + radius as f64 * t.cos(); 
//         let y = center_y as f64 + radius as f64 * t.sin();

//         let _ = image.set_pixel(x.round() as i32, y.round() as i32, color.clone());

//         raster::save(&image, "circle.png").expect("Failed to save image");

//     }
// }