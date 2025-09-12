let file = "input.txt";;
let count = ref 0
let () = 
  let ic = open_in file in
  try
    let content = really_input_string ic (in_channel_length ic) in
    String.iter (function
      | '(' -> incr count
      | ')' -> decr count
      | _ -> ()
    ) content;
    close_in ic;
    print_int !count
  with e ->
    close_in_noerr ic;
    raise e

