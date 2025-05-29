#!/bin/bash

# path Variable
CONCURRENCY_PKG=concurrency
LOAD_BALANCE_PKG=load-balancing

show_menu() {
    echo "========== MENU =========="
    echo "1) worker pool"
    echo "2) generator"
    echo "3) multiplexing"
    echo "===== Load balancing ====="
    echo "4) load balance alg: round-robin"
    echo "0) Exit"
    echo "=========================="
}

show_menu
read -p "Choose an option: " choice

case "$choice" in
1)
    clear
    echo "worker pool example:"
    go run $CONCURRENCY_PKG/worker-pool/main.go
    ;;
2)
    clear
    echo "generator example"
    go run $CONCURRENCY_PKG/generator/main.go
    ;;
3)
    clear
    echo "multiplexing example"
    go run $CONCURRENCY_PKG/multiplexing/main.go
    ;;
4)
    clear
    echo "======== APP: load balancing ========"
    echo "--------------------------------------"
    go run $LOAD_BALANCE_PKG/round-robin/main.go
    ;;
0)
    echo "Goodbye!"
    exit 0
    ;;
*)
    echo "Invalid option. Please try again."
    ;;
esac

echo "" # Add a line break
