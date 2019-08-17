#!/usr/bin/env python

import subprocess


def get_command(argvs, unit):
    cmd = []
    namespace = ""
    if len(argvs) == 2:
        namespace = argvs[1]
        cmd = ["kubectl", "get", unit, "-n", namespace]
    else:
        cmd = ["kubectl", "get", unit]
    return cmd, namespace


def span(lines, namespace):
    # Move headers two spaces to the right to account for quoting

    hf = lines[0].split(' ')
    i = 0
    begin = False
    for h in hf:
        if h == '' and not begin:
            # first space
            begin = True
            print '       ' + len(namespace) * ' ',
        elif begin and h != '':
            print h,
        elif i == 0:
            print h,
        else:
            print ' ',
    print


def format(lines, namespace, describer='kdp'):
    for l in lines[1:]:
        fields = l.split(' ')
        if len(fields) < 4:
            break
        if namespace == "":
            print '(' + describer + ' ' + fields[0] + ')',
        else:
            # Add namespace next to it, so we can double click and kdp (describe
            # it) quickly
            print '(' + describer + ' ' + fields[0] + ' ' + namespace + ')',
        i = 0
        for f in fields[1:]:
            if f == '' and i == 2:
                print ' ',
            else:
                print f,
                i = 0
            i += 1
        print


def describe(argvs, describer):
    logcmd = "kubectl logs "
    eventscmd = "kubectl get events "
    podname = ""
    namespace = ""
    if len(argvs) == 3:
        # kdp pod namespace
        # pod name + namespace
        cmd = ["kubectl", "describe", describer, argvs[1], "-n", argvs[2]]
        logcmd += argvs[1] + " -n " + argvs[2]
        eventscmd += "-n " + argvs[2]
        podname = argvs[1]
        namespace = argvs[2]
    else:
        fields = aargvs[1].split(' ')
        podname = fields[0]
        if len(fields) == 0:
            cmd = ["kubectl", "describe", describer, fields[0]]
            logcmd += fields[0]
        else:
            cmd = [
                "kubectl", "describe", describer, fields[0], "-n", fields[1]
            ]
            logcmd += fields[0] + " -n " + fields[1]
            eventscmd += "-n " + fields[1]
            namespace = fields[1]
    print subprocess.check_output(cmd)
    containers = list_containers_in_pod(podname, namespace)
    if len(containers) == 0:
        print "CHECK LOGS: (" + logcmd + ")"
    else:
        for c in containers:
            print "CHECK LOGS: (" + logcmd + " -c " + c + ")"
    print "CHECK EVENTS: (" + eventscmd + ")"
    print


def list_containers_in_pod(podname, namespace):
    cmd = "kubectl get pod " + podname + " -n " + namespace + " -o jsonpath={.spec.containers[*].name}"
    if namespace is None:
        cmd = "kubectl get pod " + podname + " -o jsonpath={.spec.containers[*].name}"
    containers = subprocess.check_output(cmd.split())
    return containers.split()
