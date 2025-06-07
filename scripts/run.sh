#!/bin/bash

# path Variable
CONCURRENCY_PKG=concurrency
LOAD_BALANCE_PKG=load-balancing

show_menu() {
    echo "========== MENU =========="
    echo "====== Concurrency ======="
    echo "1) worker pool"
    echo "2) generator"
    echo "3) multiplexing"
    echo "===== Load balancing ====="
    echo "4) alg: round-robin"
    echo "5) alg: weight-round-robin"
    echo "6) alg: source-ip-hash"
    echo "0) Exit"
    echo "=========================="
}

show_menu
read -p "Choose an option: " choice

case "$choice" in
1)
    clear
    echo "======== APP: Concurrency ========="
    echo "=========== Worker Pool ==========="
    go run $CONCURRENCY_PKG/main.go --app-name=worker-pool
    ;;
2)
    clear
    echo "======== APP: Concurrency ========="
    echo "============ Generator ============"
    go run $CONCURRENCY_PKG/main.go --app-name=generator
    ;;
3)
    clear
    echo "======== APP: Concurrency ========="
    echo "======== Multiplexing ========="
    go run $CONCURRENCY_PKG/main.go --app-name=multiplexing
    ;;
4)
    clear
    echo "======== APP: load balancing ========"
    echo "--------------------------------------"
    go run load-balancing/main.go --app-name=round-robin
    ;;
5)
    clear
    echo "======== APP: load balancing ========"
    echo "--------------------------------------"
    go run load-balancing/main.go --app-name=wrr
    ;;
6)
    clear
    echo "======== APP: load balancing ========"
    echo "--------------------------------------"
    go run load-balancing/main.go --app-name=sih
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
