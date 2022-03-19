use std::fs::File;
use std::io::{BufRead, BufReader, Result};
use serde_json::{Value, Map};

// TODO::use tokio async
fn main() -> Result<()> {
    // open file
    let file = File::open("test.log.example")?;

    // TODO::log file change monitoring

    // FIXME::the line read
    for line in BufReader::new(file).lines() {
        // parserd str to json
        let parsed: Value = serde_json::from_str(&line?)?;
        let obj: Map<String, Value> = parsed.as_object().unwrap().clone();
        println!("{:?}", obj);
    }
    Ok(())
}